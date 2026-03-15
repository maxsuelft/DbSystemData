package healthcheck_attempt

import (
	"sync"
	"sync/atomic"

	"dbsystemdata-backend/internal/features/databases"
	healthcheck_config "dbsystemdata-backend/internal/features/healthcheck/config"
	"dbsystemdata-backend/internal/features/notifiers"
	workspaces_services "dbsystemdata-backend/internal/features/workspaces/services"
	"dbsystemdata-backend/internal/util/logger"
)

var (
	healthcheckAttemptRepository = &HealthcheckAttemptRepository{}
	healthcheckAttemptService    = &HealthcheckAttemptService{
		healthcheckAttemptRepository,
		databases.GetDatabaseService(),
		workspaces_services.GetWorkspaceService(),
	}
)

var checkDatabaseHealthUseCase = &CheckDatabaseHealthUseCase{
	healthcheckAttemptRepository,
	notifiers.GetNotifierService(),
	databases.GetDatabaseService(),
}

var healthcheckAttemptBackgroundService = &HealthcheckAttemptBackgroundService{
	healthcheckConfigService:   healthcheck_config.GetHealthcheckConfigService(),
	checkDatabaseHealthUseCase: checkDatabaseHealthUseCase,
	logger:                     logger.GetLogger(),
	runOnce:                    sync.Once{},
	hasRun:                     atomic.Bool{},
}

var healthcheckAttemptController = &HealthcheckAttemptController{
	healthcheckAttemptService,
}

func GetHealthcheckAttemptRepository() *HealthcheckAttemptRepository {
	return healthcheckAttemptRepository
}

func GetHealthcheckAttemptService() *HealthcheckAttemptService {
	return healthcheckAttemptService
}

func GetHealthcheckAttemptBackgroundService() *HealthcheckAttemptBackgroundService {
	return healthcheckAttemptBackgroundService
}

func GetHealthcheckAttemptController() *HealthcheckAttemptController {
	return healthcheckAttemptController
}
