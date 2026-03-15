package usecases_mysql

import (
	"dbsystemdata-backend/internal/features/encryption/secrets"
	"dbsystemdata-backend/internal/util/encryption"
	"dbsystemdata-backend/internal/util/logger"
)

var createMysqlBackupUsecase = &CreateMysqlBackupUsecase{
	logger.GetLogger(),
	secrets.GetSecretKeyService(),
	encryption.GetFieldEncryptor(),
}

func GetCreateMysqlBackupUsecase() *CreateMysqlBackupUsecase {
	return createMysqlBackupUsecase
}
