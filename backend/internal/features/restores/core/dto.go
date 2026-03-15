package restores_core

import (
	"dbsystemdata-backend/internal/features/databases/databases/mariadb"
	"dbsystemdata-backend/internal/features/databases/databases/mongodb"
	"dbsystemdata-backend/internal/features/databases/databases/mysql"
	"dbsystemdata-backend/internal/features/databases/databases/postgresql"
)

type RestoreBackupRequest struct {
	PostgresqlDatabase *postgresql.PostgresqlDatabase `json:"postgresqlDatabase"`
	MysqlDatabase      *mysql.MysqlDatabase           `json:"mysqlDatabase"`
	MariadbDatabase    *mariadb.MariadbDatabase       `json:"mariadbDatabase"`
	MongodbDatabase    *mongodb.MongodbDatabase       `json:"mongodbDatabase"`
}
