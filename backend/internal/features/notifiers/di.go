package notifiers

import (
	"sync"
	"sync/atomic"

	audit_logs "dbsystemdata-backend/internal/features/audit_logs"
	workspaces_services "dbsystemdata-backend/internal/features/workspaces/services"
	"dbsystemdata-backend/internal/util/encryption"
	"dbsystemdata-backend/internal/util/logger"
)

var (
	notifierRepository = &NotifierRepository{}
	notifierService    = &NotifierService{
		notifierRepository,
		logger.GetLogger(),
		workspaces_services.GetWorkspaceService(),
		audit_logs.GetAuditLogService(),
		encryption.GetFieldEncryptor(),
		nil,
	}
)

var notifierController = &NotifierController{
	notifierService,
	workspaces_services.GetWorkspaceService(),
}

func GetNotifierController() *NotifierController {
	return notifierController
}

func GetNotifierService() *NotifierService {
	return notifierService
}

func GetNotifierRepository() *NotifierRepository {
	return notifierRepository
}

var (
	setupOnce sync.Once
	isSetup   atomic.Bool
)

func SetupDependencies() {
	wasAlreadySetup := isSetup.Load()

	setupOnce.Do(func() {
		workspaces_services.GetWorkspaceService().AddWorkspaceDeletionListener(notifierService)

		isSetup.Store(true)
	})

	if wasAlreadySetup {
		logger.GetLogger().Warn("SetupDependencies called multiple times, ignoring subsequent call")
	}
}
