package usecases

import (
	"context"
	"errors"

	backups_core "dbsystemdata-backend/internal/features/backups/backups/core"
	backups_config "dbsystemdata-backend/internal/features/backups/config"
	"dbsystemdata-backend/internal/features/databases"
	restores_core "dbsystemdata-backend/internal/features/restores/core"
	usecases_mariadb "dbsystemdata-backend/internal/features/restores/usecases/mariadb"
	usecases_mongodb "dbsystemdata-backend/internal/features/restores/usecases/mongodb"
	usecases_mysql "dbsystemdata-backend/internal/features/restores/usecases/mysql"
	usecases_postgresql "dbsystemdata-backend/internal/features/restores/usecases/postgresql"
	"dbsystemdata-backend/internal/features/storages"
)

type RestoreBackupUsecase struct {
	restorePostgresqlBackupUsecase *usecases_postgresql.RestorePostgresqlBackupUsecase
	restoreMysqlBackupUsecase      *usecases_mysql.RestoreMysqlBackupUsecase
	restoreMariadbBackupUsecase    *usecases_mariadb.RestoreMariadbBackupUsecase
	restoreMongodbBackupUsecase    *usecases_mongodb.RestoreMongodbBackupUsecase
}

func (uc *RestoreBackupUsecase) Execute(
	ctx context.Context,
	backupConfig *backups_config.BackupConfig,
	restore restores_core.Restore,
	originalDB *databases.Database,
	restoringToDB *databases.Database,
	backup *backups_core.Backup,
	storage *storages.Storage,
	isExcludeExtensions bool,
) error {
	switch originalDB.Type {
	case databases.DatabaseTypePostgres:
		return uc.restorePostgresqlBackupUsecase.Execute(
			ctx,
			originalDB,
			restoringToDB,
			backupConfig,
			restore,
			backup,
			storage,
			isExcludeExtensions,
		)
	case databases.DatabaseTypeMysql:
		return uc.restoreMysqlBackupUsecase.Execute(
			ctx,
			originalDB,
			restoringToDB,
			backupConfig,
			restore,
			backup,
			storage,
		)
	case databases.DatabaseTypeMariadb:
		return uc.restoreMariadbBackupUsecase.Execute(
			ctx,
			originalDB,
			restoringToDB,
			backupConfig,
			restore,
			backup,
			storage,
		)
	case databases.DatabaseTypeMongodb:
		return uc.restoreMongodbBackupUsecase.Execute(
			ctx,
			originalDB,
			restoringToDB,
			backupConfig,
			restore,
			backup,
			storage,
		)
	default:
		return errors.New("database type not supported")
	}
}
