package restores

import (
	"sync"
	"sync/atomic"

	audit_logs "dbsystemdata-backend/internal/features/audit_logs"
	"dbsystemdata-backend/internal/features/backups/backups/backuping"
	backups_services "dbsystemdata-backend/internal/features/backups/backups/services"
	backups_config "dbsystemdata-backend/internal/features/backups/config"
	"dbsystemdata-backend/internal/features/databases"
	"dbsystemdata-backend/internal/features/disk"
	restores_core "dbsystemdata-backend/internal/features/restores/core"
	"dbsystemdata-backend/internal/features/restores/usecases"
	"dbsystemdata-backend/internal/features/storages"
	tasks_cancellation "dbsystemdata-backend/internal/features/tasks/cancellation"
	workspaces_services "dbsystemdata-backend/internal/features/workspaces/services"
	"dbsystemdata-backend/internal/util/encryption"
	"dbsystemdata-backend/internal/util/logger"
)

var (
	restoreRepository = &restores_core.RestoreRepository{}
	restoreService    = &RestoreService{
		backups_services.GetBackupService(),
		restoreRepository,
		storages.GetStorageService(),
		backups_config.GetBackupConfigService(),
		usecases.GetRestoreBackupUsecase(),
		databases.GetDatabaseService(),
		logger.GetLogger(),
		workspaces_services.GetWorkspaceService(),
		audit_logs.GetAuditLogService(),
		encryption.GetFieldEncryptor(),
		disk.GetDiskService(),
		tasks_cancellation.GetTaskCancelManager(),
	}
)

var restoreController = &RestoreController{
	restoreService,
}

func GetRestoreController() *RestoreController {
	return restoreController
}

var (
	setupOnce sync.Once
	isSetup   atomic.Bool
)

func SetupDependencies() {
	wasAlreadySetup := isSetup.Load()

	setupOnce.Do(func() {
		backups_services.GetBackupService().AddBackupRemoveListener(restoreService)
		backuping.GetBackupCleaner().AddBackupRemoveListener(restoreService)

		isSetup.Store(true)
	})

	if wasAlreadySetup {
		logger.GetLogger().Warn("SetupDependencies called multiple times, ignoring subsequent call")
	}
}
