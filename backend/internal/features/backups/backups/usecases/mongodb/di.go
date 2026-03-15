package usecases_mongodb

import (
	encryption_secrets "dbsystemdata-backend/internal/features/encryption/secrets"
	"dbsystemdata-backend/internal/util/encryption"
	"dbsystemdata-backend/internal/util/logger"
)

var createMongodbBackupUsecase = &CreateMongodbBackupUsecase{
	logger.GetLogger(),
	encryption_secrets.GetSecretKeyService(),
	encryption.GetFieldEncryptor(),
}

func GetCreateMongodbBackupUsecase() *CreateMongodbBackupUsecase {
	return createMongodbBackupUsecase
}
