package dropbox_storage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"strings"
	"time"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
	"github.com/google/uuid"
	"golang.org/x/oauth2"

	"dbsystemdata-backend/internal/util/encryption"
)

const (
	// dropboxBackupsFolder is the root folder for all backups. Each database uses a subfolder
	// so the structure is: /dbsystemdata_backups/<database_name>/<backup_file>.
	// Backup FileName is generated as "databaseName/databaseName-timestamp-id.ext" in backup core.
	dropboxBackupsFolder = "/dbsystemdata_backups"
	dropboxTimeout       = 60 * time.Second
)

// buildStoragePath returns the full Dropbox path: /dbsystemdata_backups/<fileName>.
// For backups, fileName is "databaseName/databaseName-timestamp-id.ext", so files end up
// in subfolders per database: dbsystemdata_backups/nomedobancodedado/...
func buildStoragePath(fileName string) string {
	fileName = strings.TrimPrefix(fileName, "/")
	if fileName == "" {
		return dropboxBackupsFolder
	}
	return dropboxBackupsFolder + "/" + fileName
}

var dropboxOAuth2Endpoint = oauth2.Endpoint{
	AuthURL:  "https://www.dropbox.com/oauth2/authorize",
	TokenURL: "https://api.dropboxapi.com/oauth2/token",
}

type DropboxStorage struct {
	StorageID  uuid.UUID `json:"storageId"  gorm:"primaryKey;type:uuid;column:storage_id"`
	AppKey     string    `json:"appKey"     gorm:"not null;type:text;column:app_key"`
	AppSecret  string    `json:"appSecret"  gorm:"not null;type:text;column:app_secret"`
	TokenJSON  string    `json:"tokenJson"  gorm:"not null;type:text;column:token_json"`
}

func (s *DropboxStorage) TableName() string {
	return "dropbox_storages"
}

func (s *DropboxStorage) SaveFile(
	ctx context.Context,
	encryptor encryption.FieldEncryptor,
	logger *slog.Logger,
	fileName string,
	file io.Reader,
) error {
	client, err := s.getClient(encryptor)
	if err != nil {
		return err
	}

	path := buildStoragePath(fileName)

	ctx, cancel := context.WithTimeout(ctx, dropboxTimeout)
	defer cancel()

	arg := files.NewUploadArg(path)
	arg.Mode = &files.WriteMode{Tagged: dropbox.Tagged{Tag: files.WriteModeOverwrite}}
	_, err = client.Upload(arg, file)
	if err != nil {
		return fmt.Errorf("failed to upload file to Dropbox: %w", err)
	}

	logger.Info("file uploaded to Dropbox", "name", fileName, "path", path)
	return nil
}

func (s *DropboxStorage) GetFile(
	encryptor encryption.FieldEncryptor,
	fileName string,
) (io.ReadCloser, error) {
	client, err := s.getClient(encryptor)
	if err != nil {
		return nil, err
	}

	path := buildStoragePath(fileName)

	_, content, err := client.Download(files.NewDownloadArg(path))
	if err != nil {
		return nil, fmt.Errorf("failed to download file from Dropbox: %w", err)
	}

	return content, nil
}

func (s *DropboxStorage) DeleteFile(
	encryptor encryption.FieldEncryptor,
	fileName string,
) error {
	client, err := s.getClient(encryptor)
	if err != nil {
		return err
	}

	path := buildStoragePath(fileName)

	_, err = client.DeleteV2(files.NewDeleteArg(path))
	if err != nil {
		return fmt.Errorf("failed to delete file from Dropbox: %w", err)
	}

	return nil
}

func (s *DropboxStorage) Validate(encryptor encryption.FieldEncryptor) error {
	switch {
	case s.AppKey == "":
		return errors.New("app key is required")
	case s.AppSecret == "":
		return errors.New("app secret is required")
	case s.TokenJSON == "":
		return errors.New("token JSON is required")
	}

	if strings.HasPrefix(s.TokenJSON, "enc:") {
		return nil
	}

	var token oauth2.Token
	if err := json.Unmarshal([]byte(s.TokenJSON), &token); err != nil {
		return fmt.Errorf("invalid token JSON format: %w", err)
	}

	if token.RefreshToken == "" && token.AccessToken == "" {
		return errors.New("token JSON must contain access token or refresh token")
	}

	return nil
}

func (s *DropboxStorage) TestConnection(encryptor encryption.FieldEncryptor) error {
	client, err := s.getClient(encryptor)
	if err != nil {
		return err
	}

	testPath := buildStoragePath("test-connection-" + uuid.New().String())
	testContent := strings.NewReader("test")

	testArg := files.NewUploadArg(testPath)
	testArg.Mode = &files.WriteMode{Tagged: dropbox.Tagged{Tag: files.WriteModeAdd}}
	_, err = client.Upload(testArg, testContent)
	if err != nil {
		return fmt.Errorf("failed to write test file to Dropbox: %w", err)
	}

	_, err = client.DeleteV2(files.NewDeleteArg(testPath))
	if err != nil {
		return fmt.Errorf("failed to delete test file from Dropbox: %w", err)
	}

	return nil
}

func (s *DropboxStorage) HideSensitiveData() {
	s.AppSecret = ""
	s.TokenJSON = ""
}

func (s *DropboxStorage) EncryptSensitiveData(encryptor encryption.FieldEncryptor) error {
	var err error
	if s.AppSecret != "" {
		s.AppSecret, err = encryptor.Encrypt(s.StorageID, s.AppSecret)
		if err != nil {
			return fmt.Errorf("failed to encrypt Dropbox app secret: %w", err)
		}
	}
	if s.TokenJSON != "" {
		s.TokenJSON, err = encryptor.Encrypt(s.StorageID, s.TokenJSON)
		if err != nil {
			return fmt.Errorf("failed to encrypt Dropbox token JSON: %w", err)
		}
	}
	return nil
}

func (s *DropboxStorage) Update(incoming *DropboxStorage) {
	s.AppKey = incoming.AppKey
	if incoming.AppSecret != "" {
		s.AppSecret = incoming.AppSecret
	}
	if incoming.TokenJSON != "" {
		s.TokenJSON = incoming.TokenJSON
	}
}

func (s *DropboxStorage) getClient(encryptor encryption.FieldEncryptor) (files.Client, error) {
	if err := s.Validate(encryptor); err != nil {
		return nil, err
	}

	appSecret, err := encryptor.Decrypt(s.StorageID, s.AppSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt app secret: %w", err)
	}

	tokenJSON, err := encryptor.Decrypt(s.StorageID, s.TokenJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt token JSON: %w", err)
	}

	var token oauth2.Token
	if err := json.Unmarshal([]byte(tokenJSON), &token); err != nil {
		return nil, fmt.Errorf("invalid token JSON: %w", err)
	}

	cfg := &oauth2.Config{
		ClientID:     s.AppKey,
		ClientSecret: appSecret,
		Endpoint:     dropboxOAuth2Endpoint,
	}

	tokenSource := cfg.TokenSource(context.Background(), &token)
	accessToken, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get Dropbox access token: %w", err)
	}

	dbxConfig := dropbox.Config{
		Token: accessToken.AccessToken,
	}

	client := files.New(dbxConfig)
	return client, nil
}
