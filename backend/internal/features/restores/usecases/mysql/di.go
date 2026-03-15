package usecases_mysql

import (
	"dbsystemdata-backend/internal/features/encryption/secrets"
	"dbsystemdata-backend/internal/util/logger"
)

var restoreMysqlBackupUsecase = &RestoreMysqlBackupUsecase{
	logger.GetLogger(),
	secrets.GetSecretKeyService(),
}

func GetRestoreMysqlBackupUsecase() *RestoreMysqlBackupUsecase {
	return restoreMysqlBackupUsecase
}
