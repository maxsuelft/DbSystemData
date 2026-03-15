package restoring

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"

	backups_services "dbsystemdata-backend/internal/features/backups/backups/services"
	backups_config "dbsystemdata-backend/internal/features/backups/config"
	"dbsystemdata-backend/internal/features/databases"
	restores_core "dbsystemdata-backend/internal/features/restores/core"
	"dbsystemdata-backend/internal/features/restores/usecases"
	"dbsystemdata-backend/internal/features/storages"
	tasks_cancellation "dbsystemdata-backend/internal/features/tasks/cancellation"
	cache_utils "dbsystemdata-backend/internal/util/cache"
	"dbsystemdata-backend/internal/util/encryption"
	"dbsystemdata-backend/internal/util/logger"
)

var restoreRepository = &restores_core.RestoreRepository{}

var restoreNodesRegistry = &RestoreNodesRegistry{
	client:            cache_utils.GetValkeyClient(),
	logger:            logger.GetLogger(),
	timeout:           cache_utils.DefaultCacheTimeout,
	pubsubRestores:    cache_utils.NewPubSubManager(),
	pubsubCompletions: cache_utils.NewPubSubManager(),
	runOnce:           sync.Once{},
	hasRun:            atomic.Bool{},
}

var restoreDatabaseCache = cache_utils.NewCacheUtil[RestoreDatabaseCache](
	cache_utils.GetValkeyClient(),
	"restore_db:",
)

var restoreCancelManager = tasks_cancellation.GetTaskCancelManager()

var restorerNode = &RestorerNode{
	uuid.New(),
	databases.GetDatabaseService(),
	backups_services.GetBackupService(),
	encryption.GetFieldEncryptor(),
	restoreRepository,
	backups_config.GetBackupConfigService(),
	storages.GetStorageService(),
	restoreNodesRegistry,
	logger.GetLogger(),
	usecases.GetRestoreBackupUsecase(),
	restoreDatabaseCache,
	restoreCancelManager,
	time.Time{},
	sync.Once{},
	atomic.Bool{},
}

var restoresScheduler = &RestoresScheduler{
	restoreRepository,
	backups_services.GetBackupService(),
	storages.GetStorageService(),
	backups_config.GetBackupConfigService(),
	restoreNodesRegistry,
	time.Now().UTC(),
	logger.GetLogger(),
	make(map[uuid.UUID]RestoreToNodeRelation),
	restorerNode,
	restoreDatabaseCache,
	uuid.Nil,
	sync.Once{},
	atomic.Bool{},
}

func GetRestoresScheduler() *RestoresScheduler {
	return restoresScheduler
}

func GetRestorerNode() *RestorerNode {
	return restorerNode
}

func GetRestoreNodesRegistry() *RestoreNodesRegistry {
	return restoreNodesRegistry
}
