package backups_download

import (
	"sync"
	"sync/atomic"

	"dbsystemdata-backend/internal/config"
	cache_utils "dbsystemdata-backend/internal/util/cache"
	"dbsystemdata-backend/internal/util/logger"
)

var downloadTokenRepository = &DownloadTokenRepository{}

var downloadTracker = NewDownloadTracker(cache_utils.GetValkeyClient())

var (
	bandwidthManager               *BandwidthManager
	downloadTokenService           *DownloadTokenService
	downloadTokenBackgroundService *DownloadTokenBackgroundService
)

func init() {
	env := config.GetEnv()
	throughputMBs := env.NodeNetworkThroughputMBs
	if throughputMBs == 0 {
		throughputMBs = 125
	}
	bandwidthManager = NewBandwidthManager(throughputMBs)

	downloadTokenService = &DownloadTokenService{
		downloadTokenRepository,
		logger.GetLogger(),
		downloadTracker,
		bandwidthManager,
	}

	downloadTokenBackgroundService = &DownloadTokenBackgroundService{
		downloadTokenService: downloadTokenService,
		logger:               logger.GetLogger(),
		runOnce:              sync.Once{},
		hasRun:               atomic.Bool{},
	}
}

func GetDownloadTokenService() *DownloadTokenService {
	return downloadTokenService
}

func GetDownloadTokenBackgroundService() *DownloadTokenBackgroundService {
	return downloadTokenBackgroundService
}

func GetBandwidthManager() *BandwidthManager {
	return bandwidthManager
}
