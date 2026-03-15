package backuping

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"

	backups_core "dbsystemdata-backend/internal/features/backups/backups/core"
	"dbsystemdata-backend/internal/features/backups/backups/usecases"
	backups_config "dbsystemdata-backend/internal/features/backups/config"
	"dbsystemdata-backend/internal/features/databases"
	"dbsystemdata-backend/internal/features/notifiers"
	"dbsystemdata-backend/internal/features/storages"
	tasks_cancellation "dbsystemdata-backend/internal/features/tasks/cancellation"
	workspaces_services "dbsystemdata-backend/internal/features/workspaces/services"
	cache_utils "dbsystemdata-backend/internal/util/cache"
	"dbsystemdata-backend/internal/util/encryption"
	"dbsystemdata-backend/internal/util/logger"
)

var backupRepository = &backups_core.BackupRepository{}

var taskCancelManager = tasks_cancellation.GetTaskCancelManager()

var backupCleaner = &BackupCleaner{
	backupRepository,
	storages.GetStorageService(),
	backups_config.GetBackupConfigService(),
	encryption.GetFieldEncryptor(),
	logger.GetLogger(),
	[]backups_core.BackupRemoveListener{},
	sync.Once{},
	atomic.Bool{},
}

var backupNodesRegistry = &BackupNodesRegistry{
	cache_utils.GetValkeyClient(),
	logger.GetLogger(),
	cache_utils.DefaultCacheTimeout,
	cache_utils.NewPubSubManager(),
	cache_utils.NewPubSubManager(),
	sync.Once{},
	atomic.Bool{},
}

func getNodeID() uuid.UUID {
	return uuid.New()
}

var backuperNode = &BackuperNode{
	databases.GetDatabaseService(),
	encryption.GetFieldEncryptor(),
	workspaces_services.GetWorkspaceService(),
	backupRepository,
	backups_config.GetBackupConfigService(),
	storages.GetStorageService(),
	notifiers.GetNotifierService(),
	taskCancelManager,
	backupNodesRegistry,
	logger.GetLogger(),
	usecases.GetCreateBackupUsecase(),
	getNodeID(),
	time.Time{},
	sync.Once{},
	atomic.Bool{},
}

var backupsScheduler = &BackupsScheduler{
	backupRepository,
	backups_config.GetBackupConfigService(),
	taskCancelManager,
	backupNodesRegistry,
	databases.GetDatabaseService(),
	time.Now().UTC(),
	logger.GetLogger(),
	make(map[uuid.UUID]BackupToNodeRelation),
	backuperNode,
	sync.Once{},
	atomic.Bool{},
}

func GetBackupsScheduler() *BackupsScheduler {
	return backupsScheduler
}

func GetBackuperNode() *BackuperNode {
	return backuperNode
}

func GetBackupNodesRegistry() *BackupNodesRegistry {
	return backupNodesRegistry
}

func GetBackupCleaner() *BackupCleaner {
	return backupCleaner
}
