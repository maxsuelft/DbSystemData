package backups_core

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	backups_config "dbsystemdata-backend/internal/features/backups/config"
	files_utils "dbsystemdata-backend/internal/util/files"
)

type PgWalBackupType string

const (
	PgWalBackupTypeFullBackup PgWalBackupType = "PG_FULL_BACKUP"
	PgWalBackupTypeWalSegment PgWalBackupType = "PG_WAL_SEGMENT"
)

type Backup struct {
	ID       uuid.UUID `json:"id"       gorm:"column:id;type:uuid;primaryKey"`
	FileName string    `json:"fileName" gorm:"column:file_name;type:text;not null"`

	DatabaseID uuid.UUID `json:"databaseId" gorm:"column:database_id;type:uuid;not null"`
	StorageID  uuid.UUID `json:"storageId"  gorm:"column:storage_id;type:uuid;not null"`

	Status      BackupStatus `json:"status"      gorm:"column:status;not null"`
	FailMessage *string      `json:"failMessage" gorm:"column:fail_message"`
	IsSkipRetry bool         `json:"isSkipRetry" gorm:"column:is_skip_retry;type:boolean;not null"`

	BackupSizeMb float64 `json:"backupSizeMb" gorm:"column:backup_size_mb;default:0"`

	BackupDurationMs int64 `json:"backupDurationMs" gorm:"column:backup_duration_ms;default:0"`

	EncryptionSalt *string                         `json:"-"          gorm:"column:encryption_salt"`
	EncryptionIV   *string                         `json:"-"          gorm:"column:encryption_iv"`
	Encryption     backups_config.BackupEncryption `json:"encryption" gorm:"column:encryption;type:text;not null;default:'NONE'"`

	// Postgres WAL backup specific fields
	PgWalBackupType                 *PgWalBackupType `json:"pgWalBackupType"                 gorm:"column:pg_wal_backup_type;type:text"`
	PgFullBackupWalStartSegmentName *string          `json:"pgFullBackupWalStartSegmentName" gorm:"column:pg_wal_start_segment;type:text"`
	PgFullBackupWalStopSegmentName  *string          `json:"pgFullBackupWalStopSegmentName"  gorm:"column:pg_wal_stop_segment;type:text"`
	PgVersion                       *string          `json:"pgVersion"                       gorm:"column:pg_version;type:text"`
	PgWalSegmentName                *string          `json:"pgWalSegmentName"                gorm:"column:pg_wal_segment_name;type:text"`

	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
}

// GenerateFilename sets the backup file name. When extension is non-empty (e.g. when
// encryption is None), the file is stored with that extension so it keeps its real format.
// When databaseID is non-nil, the path is prefixed with databaseID so backups are stored
// in a subfolder per database (e.g. dbsystemdata_backups/<databaseID>/file.dump).
func (b *Backup) GenerateFilename(dbName string, extension string, databaseID *uuid.UUID) {
	timestamp := time.Now().UTC()

	base := fmt.Sprintf(
		"%s-%s-%s%s",
		files_utils.SanitizeFilename(dbName),
		timestamp.Format("20060102-150405"),
		b.ID.String(),
		extension,
	)
	if databaseID != nil {
		b.FileName = databaseID.String() + "/" + base
	} else {
		b.FileName = base
	}
}
