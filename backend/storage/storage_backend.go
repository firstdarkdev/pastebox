package storage

// StorageBackend - Storage Backend Interface
type StorageBackend interface {
	Save(key string, data string) (string, error)
	Read(key string) (string, error)
}
