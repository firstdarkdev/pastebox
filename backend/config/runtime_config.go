package config

import (
	"haste/keygen"
	"haste/storage"
)

// RuntimeConfig Config used throughout the application
type RuntimeConfig struct {
	Generator keygen.KeyGenerator
	Storage   *storage.FileStorage
	Conf      Config
}
