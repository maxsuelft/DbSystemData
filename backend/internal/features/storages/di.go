package storages

import (
	"sync"
	"sync/atomic"

	audit_logs "dbsystemdata-backend/internal/features/audit_logs"
	workspaces_services "dbsystemdata-backend/internal/features/workspaces/services"
	"dbsystemdata-backend/internal/util/encryption"
	"dbsystemdata-backend/internal/util/logger"
)

var (
	storageRepository = &StorageRepository{}
	storageService    = &StorageService{
		storageRepository,
		workspaces_services.GetWorkspaceService(),
		audit_logs.GetAuditLogService(),
		encryption.GetFieldEncryptor(),
		nil,
	}
)

var storageController = &StorageController{
	storageService,
	workspaces_services.GetWorkspaceService(),
}

func GetStorageService() *StorageService {
	return storageService
}

func GetStorageController() *StorageController {
	return storageController
}

var (
	setupOnce sync.Once
	isSetup   atomic.Bool
)

func SetupDependencies() {
	wasAlreadySetup := isSetup.Load()

	setupOnce.Do(func() {
		workspaces_services.GetWorkspaceService().AddWorkspaceDeletionListener(storageService)

		isSetup.Store(true)
	})

	if wasAlreadySetup {
		logger.GetLogger().Warn("SetupDependencies called multiple times, ignoring subsequent call")
	}
}
