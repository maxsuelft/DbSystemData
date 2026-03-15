package usecases_mariadb

import (
	"dbsystemdata-backend/internal/features/encryption/secrets"
	"dbsystemdata-backend/internal/util/logger"
)

var restoreMariadbBackupUsecase = &RestoreMariadbBackupUsecase{
	logger.GetLogger(),
	secrets.GetSecretKeyService(),
}

func GetRestoreMariadbBackupUsecase() *RestoreMariadbBackupUsecase {
	return restoreMariadbBackupUsecase
}
