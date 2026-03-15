package restoring

import (
	"time"

	"github.com/google/uuid"

	"dbsystemdata-backend/internal/features/databases/databases/mariadb"
	"dbsystemdata-backend/internal/features/databases/databases/mongodb"
	"dbsystemdata-backend/internal/features/databases/databases/mysql"
	"dbsystemdata-backend/internal/features/databases/databases/postgresql"
)

type RestoreDatabaseCache struct {
	PostgresqlDatabase *postgresql.PostgresqlDatabase `json:"postgresqlDatabase,omitempty"`
	MysqlDatabase      *mysql.MysqlDatabase           `json:"mysqlDatabase,omitempty"`
	MariadbDatabase    *mariadb.MariadbDatabase       `json:"mariadbDatabase,omitempty"`
	MongodbDatabase    *mongodb.MongodbDatabase       `json:"mongodbDatabase,omitempty"`
}

type RestoreToNodeRelation struct {
	NodeID     uuid.UUID   `json:"nodeId"`
	RestoreIDs []uuid.UUID `json:"restoreIds"`
}

type RestoreNode struct {
	ID            uuid.UUID `json:"id"`
	ThroughputMBs int       `json:"throughputMBs"`
	LastHeartbeat time.Time `json:"lastHeartbeat"`
}

type RestoreNodeStats struct {
	ID             uuid.UUID `json:"id"`
	ActiveRestores int       `json:"activeRestores"`
}

type RestoreSubmitMessage struct {
	NodeID         uuid.UUID `json:"nodeId"`
	RestoreID      uuid.UUID `json:"restoreId"`
	IsCallNotifier bool      `json:"isCallNotifier"`
}

type RestoreCompletionMessage struct {
	NodeID    uuid.UUID `json:"nodeId"`
	RestoreID uuid.UUID `json:"restoreId"`
}
