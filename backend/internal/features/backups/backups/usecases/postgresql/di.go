package usecases_postgresql

import (
	"dbsystemdata-backend/internal/features/encryption/secrets"
	"dbsystemdata-backend/internal/util/encryption"
	"dbsystemdata-backend/internal/util/logger"
)

var createPostgresqlBackupUsecase = &CreatePostgresqlBackupUsecase{
	logger.GetLogger(),
	secrets.GetSecretKeyService(),
	encryption.GetFieldEncryptor(),
}

func GetCreatePostgresqlBackupUsecase() *CreatePostgresqlBackupUsecase {
	return createPostgresqlBackupUsecase
}
