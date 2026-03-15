package backups_core

import (
	"context"

	usecases_common "dbsystemdata-backend/internal/features/backups/backups/common"
	backups_config "dbsystemdata-backend/internal/features/backups/config"
	"dbsystemdata-backend/internal/features/databases"
	"dbsystemdata-backend/internal/features/notifiers"
	"dbsystemdata-backend/internal/features/storages"
)

type NotificationSender interface {
	SendNotification(
		notifier *notifiers.Notifier,
		title string,
		message string,
	)
}

type CreateBackupUsecase interface {
	Execute(
		ctx context.Context,
		backup *Backup,
		backupConfig *backups_config.BackupConfig,
		database *databases.Database,
		storage *storages.Storage,
		backupProgressListener func(completedMBs float64),
	) (*usecases_common.BackupMetadata, error)
}

type BackupRemoveListener interface {
	OnBeforeBackupRemove(backup *Backup) error
}
