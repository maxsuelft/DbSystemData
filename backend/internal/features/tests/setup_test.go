package tests

import (
	"os"
	"testing"

	"dbsystemdata-backend/internal/features/backups/backups/backuping"
	"dbsystemdata-backend/internal/features/restores/restoring"
	cache_utils "dbsystemdata-backend/internal/util/cache"
)

func TestMain(m *testing.M) {
	cache_utils.ClearAllCache()

	backuperNode := backuping.CreateTestBackuperNode()
	cancelBackup := backuping.StartBackuperNodeForTest(&testing.T{}, backuperNode)

	restorerNode := restoring.CreateTestRestorerNode()
	cancelRestore := restoring.StartRestorerNodeForTest(&testing.T{}, restorerNode)

	exitCode := m.Run()

	backuping.StopBackuperNodeForTest(&testing.T{}, cancelBackup, backuperNode)
	restoring.StopRestorerNodeForTest(&testing.T{}, cancelRestore, restorerNode)

	os.Exit(exitCode)
}
