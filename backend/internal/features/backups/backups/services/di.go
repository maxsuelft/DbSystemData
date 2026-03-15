package backups_services

import (
	"sync"
	"sync/atomic"

	audit_logs "dbsystemdata-backend/internal/features/audit_logs"
	"dbsystemdata-backend/internal/features/backups/backups/backuping"
	backups_core "dbsystemdata-backend/internal/features/backups/backups/core"
	backups_download "dbsystemdata-backend/internal/features/backups/backups/download"
	"dbsystemdata-backend/internal/features/backups/backups/usecases"
	backups_config "dbsystemdata-backend/internal/features/backups/config"
	"dbsystemdata-backend/internal/features/databases"
	encryption_secrets "dbsystemdata-backend/internal/features/encryption/secrets"
	"dbsystemdata-backend/internal/features/notifiers"
	"dbsystemdata-backend/internal/features/storages"
	task_cancellation "dbsystemdata-backend/internal/features/tasks/cancellation"
	workspaces_services "dbsystemdata-backend/internal/features/workspaces/services"
	"dbsystemdata-backend/internal/util/encryption"
	"dbsystemdata-backend/internal/util/logger"
)

var taskCancelManager = task_cancellation.GetTaskCancelManager()

var backupService = &BackupService{
	databases.GetDatabaseService(),
	storages.GetStorageService(),
	backups_core.GetBackupRepository(),
	notifiers.GetNotifierService(),
	notifiers.GetNotifierService(),
	backups_config.GetBackupConfigService(),
	encryption_secrets.GetSecretKeyService(),
	encryption.GetFieldEncryptor(),
	usecases.GetCreateBackupUsecase(),
	logger.GetLogger(),
	[]backups_core.BackupRemoveListener{},
	workspaces_services.GetWorkspaceService(),
	audit_logs.GetAuditLogService(),
	taskCancelManager,
	backups_download.GetDownloadTokenService(),
	backuping.GetBackupsScheduler(),
	backuping.GetBackupCleaner(),
}

func GetBackupService() *BackupService {
	return backupService
}

var walService = &PostgreWalBackupService{
	backups_config.GetBackupConfigService(),
	backups_core.GetBackupRepository(),
	encryption.GetFieldEncryptor(),
	encryption_secrets.GetSecretKeyService(),
	logger.GetLogger(),
	backupService,
}

func GetWalService() *PostgreWalBackupService {
	return walService
}

var (
	setupOnce sync.Once
	isSetup   atomic.Bool
)

func SetupDependencies() {
	wasAlreadySetup := isSetup.Load()

	setupOnce.Do(func() {
		backups_config.
			GetBackupConfigService().
			SetDatabaseStorageChangeListener(backupService)

		databases.GetDatabaseService().AddDbRemoveListener(backupService)
		databases.GetDatabaseService().AddDbCopyListener(backups_config.GetBackupConfigService())

		isSetup.Store(true)
	})

	if wasAlreadySetup {
		logger.GetLogger().Warn("SetupDependencies called multiple times, ignoring subsequent call")
	}
}
