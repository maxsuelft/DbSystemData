package plans

import (
	"dbsystemdata-backend/internal/util/logger"
)

var databasePlanRepository = &DatabasePlanRepository{}

var databasePlanService = &DatabasePlanService{
	databasePlanRepository,
	logger.GetLogger(),
}

func GetDatabasePlanService() *DatabasePlanService {
	return databasePlanService
}

func GetDatabasePlanRepository() *DatabasePlanRepository {
	return databasePlanRepository
}
