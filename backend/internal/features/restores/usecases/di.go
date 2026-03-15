package usecases

import (
	usecases_mariadb "dbsystemdata-backend/internal/features/restores/usecases/mariadb"
	usecases_mongodb "dbsystemdata-backend/internal/features/restores/usecases/mongodb"
	usecases_mysql "dbsystemdata-backend/internal/features/restores/usecases/mysql"
	usecases_postgresql "dbsystemdata-backend/internal/features/restores/usecases/postgresql"
)

var restoreBackupUsecase = &RestoreBackupUsecase{
	usecases_postgresql.GetRestorePostgresqlBackupUsecase(),
	usecases_mysql.GetRestoreMysqlBackupUsecase(),
	usecases_mariadb.GetRestoreMariadbBackupUsecase(),
	usecases_mongodb.GetRestoreMongodbBackupUsecase(),
}

func GetRestoreBackupUsecase() *RestoreBackupUsecase {
	return restoreBackupUsecase
}
