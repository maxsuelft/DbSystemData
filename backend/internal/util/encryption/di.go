package encryption

import "dbsystemdata-backend/internal/features/encryption/secrets"

var fieldEncryptor = &SecretKeyFieldEncryptor{
	secrets.GetSecretKeyService(),
}

func GetFieldEncryptor() FieldEncryptor {
	return fieldEncryptor
}
