package usecases_postgresql

import (
	"dbsystemdata-backend/internal/features/encryption/secrets"
	"dbsystemdata-backend/internal/util/logger"
)

var restorePostgresqlBackupUsecase = &RestorePostgresqlBackupUsecase{
	logger.GetLogger(),
	secrets.GetSecretKeyService(),
}

func GetRestorePostgresqlBackupUsecase() *RestorePostgresqlBackupUsecase {
	return restorePostgresqlBackupUsecase
}
