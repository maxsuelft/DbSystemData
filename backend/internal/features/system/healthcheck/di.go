package system_healthcheck

import (
	"dbsystemdata-backend/internal/features/backups/backups/backuping"
	"dbsystemdata-backend/internal/features/disk"
)

var healthcheckService = &HealthcheckService{
	disk.GetDiskService(),
	backuping.GetBackupsScheduler(),
	backuping.GetBackuperNode(),
}

var healthcheckController = &HealthcheckController{
	healthcheckService,
}

func GetHealthcheckController() *HealthcheckController {
	return healthcheckController
}
