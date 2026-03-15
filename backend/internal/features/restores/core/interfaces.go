package restores_core

import (
	"context"

	backups_core "dbsystemdata-backend/internal/features/backups/backups/core"
	backups_config "dbsystemdata-backend/internal/features/backups/config"
	"dbsystemdata-backend/internal/features/databases"
	"dbsystemdata-backend/internal/features/storages"
)

type RestoreBackupUsecase interface {
	Execute(
		ctx context.Context,
		backupConfig *backups_config.BackupConfig,
		restore Restore,
		originalDB *databases.Database,
		restoringToDB *databases.Database,
		backup *backups_core.Backup,
		storage *storages.Storage,
		isExcludeExtensions bool,
	) error
}
