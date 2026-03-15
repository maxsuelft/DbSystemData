package backups_config

import (
	"sync"
	"sync/atomic"

	"dbsystemdata-backend/internal/features/databases"
	"dbsystemdata-backend/internal/features/notifiers"
	plans "dbsystemdata-backend/internal/features/plan"
	"dbsystemdata-backend/internal/features/storages"
	workspaces_services "dbsystemdata-backend/internal/features/workspaces/services"
	"dbsystemdata-backend/internal/util/logger"
)

var (
	backupConfigRepository = &BackupConfigRepository{}
	backupConfigService    = &BackupConfigService{
		backupConfigRepository,
		databases.GetDatabaseService(),
		storages.GetStorageService(),
		notifiers.GetNotifierService(),
		workspaces_services.GetWorkspaceService(),
		plans.GetDatabasePlanService(),
		nil,
	}
)

var backupConfigController = &BackupConfigController{
	backupConfigService,
}

func GetBackupConfigController() *BackupConfigController {
	return backupConfigController
}

func GetBackupConfigService() *BackupConfigService {
	return backupConfigService
}

var (
	setupOnce sync.Once
	isSetup   atomic.Bool
)

func SetupDependencies() {
	wasAlreadySetup := isSetup.Load()

	setupOnce.Do(func() {
		storages.GetStorageService().SetStorageDatabaseCounter(backupConfigService)

		isSetup.Store(true)
	})

	if wasAlreadySetup {
		logger.GetLogger().Warn("SetupDependencies called multiple times, ignoring subsequent call")
	}
}
