package usecases

import (
	"context"
	"errors"

	common "dbsystemdata-backend/internal/features/backups/backups/common"
	backups_core "dbsystemdata-backend/internal/features/backups/backups/core"
	usecases_mariadb "dbsystemdata-backend/internal/features/backups/backups/usecases/mariadb"
	usecases_mongodb "dbsystemdata-backend/internal/features/backups/backups/usecases/mongodb"
	usecases_mysql "dbsystemdata-backend/internal/features/backups/backups/usecases/mysql"
	usecases_postgresql "dbsystemdata-backend/internal/features/backups/backups/usecases/postgresql"
	backups_config "dbsystemdata-backend/internal/features/backups/config"
	"dbsystemdata-backend/internal/features/databases"
	"dbsystemdata-backend/internal/features/storages"
)

type CreateBackupUsecase struct {
	CreatePostgresqlBackupUsecase *usecases_postgresql.CreatePostgresqlBackupUsecase
	CreateMysqlBackupUsecase      *usecases_mysql.CreateMysqlBackupUsecase
	CreateMariadbBackupUsecase    *usecases_mariadb.CreateMariadbBackupUsecase
	CreateMongodbBackupUsecase    *usecases_mongodb.CreateMongodbBackupUsecase
}

func (uc *CreateBackupUsecase) Execute(
	ctx context.Context,
	backup *backups_core.Backup,
	backupConfig *backups_config.BackupConfig,
	database *databases.Database,
	storage *storages.Storage,
	backupProgressListener func(completedMBs float64),
) (*common.BackupMetadata, error) {
	switch database.Type {
	case databases.DatabaseTypePostgres:
		return uc.CreatePostgresqlBackupUsecase.Execute(
			ctx,
			backup,
			backupConfig,
			database,
			storage,
			backupProgressListener,
		)

	case databases.DatabaseTypeMysql:
		return uc.CreateMysqlBackupUsecase.Execute(
			ctx,
			backup,
			backupConfig,
			database,
			storage,
			backupProgressListener,
		)

	case databases.DatabaseTypeMariadb:
		return uc.CreateMariadbBackupUsecase.Execute(
			ctx,
			backup,
			backupConfig,
			database,
			storage,
			backupProgressListener,
		)

	case databases.DatabaseTypeMongodb:
		return uc.CreateMongodbBackupUsecase.Execute(
			ctx,
			backup,
			backupConfig,
			database,
			storage,
			backupProgressListener,
		)

	default:
		return nil, errors.New("database type not supported")
	}
}
