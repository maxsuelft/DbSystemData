package backups_config

import (
	"dbsystemdata-backend/internal/features/databases"
)

// GetStorageBackupFileExtension returns the file extension for a backup file when saved to storage.
// When encryption is None, the file is stored in its native format and gets the appropriate
// extension (.dump, .archive, .sql.zst). When encryption is Encrypted, the file is stored
// as an opaque blob with no extension.
func GetStorageBackupFileExtension(dbType databases.DatabaseType, encryption BackupEncryption) string {
	if encryption != BackupEncryptionNone {
		return ""
	}
	switch dbType {
	case databases.DatabaseTypeMysql, databases.DatabaseTypeMariadb:
		return ".sql.zst"
	case databases.DatabaseTypePostgres:
		return ".dump"
	case databases.DatabaseTypeMongodb:
		return ".archive"
	default:
		return ".backup"
	}
}
