package usecases_mariadb

import (
	"dbsystemdata-backend/internal/features/encryption/secrets"
	"dbsystemdata-backend/internal/util/encryption"
	"dbsystemdata-backend/internal/util/logger"
)

var createMariadbBackupUsecase = &CreateMariadbBackupUsecase{
	logger.GetLogger(),
	secrets.GetSecretKeyService(),
	encryption.GetFieldEncryptor(),
}

func GetCreateMariadbBackupUsecase() *CreateMariadbBackupUsecase {
	return createMariadbBackupUsecase
}
