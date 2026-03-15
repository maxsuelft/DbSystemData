package usecases_mongodb

import (
	encryption_secrets "dbsystemdata-backend/internal/features/encryption/secrets"
	"dbsystemdata-backend/internal/util/logger"
)

var restoreMongodbBackupUsecase = &RestoreMongodbBackupUsecase{
	logger.GetLogger(),
	encryption_secrets.GetSecretKeyService(),
}

func GetRestoreMongodbBackupUsecase() *RestoreMongodbBackupUsecase {
	return restoreMongodbBackupUsecase
}
