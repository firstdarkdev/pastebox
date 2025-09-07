package storage

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

const storageDir = "./data"

// FileStorage - File Storage Backend
type FileStorage struct{}

func NewFileStorage() *FileStorage {
	log.Debug("Creating Storage Directory")
	if _, err := os.Stat(storageDir); os.IsNotExist(err) {
		if err := os.MkdirAll(storageDir, 0755); err != nil {
			log.Fatal(err)
		}
	}

	return &FileStorage{}
}

func (f *FileStorage) Save(key string, data []byte) (string, error) {
	filePath := fmt.Sprintf("%s/%s", storageDir, key)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.WriteFile(filePath, data, 0644); err != nil {
			log.Error(err)
			return "", err
		}
	}

	return key, nil
}

func (f *FileStorage) Read(key string) (string, error) {
	filePath := fmt.Sprintf("%s/%s", storageDir, key)
	log.Debugf("Retrieving raw file %s", filePath)

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Error("Failed to load file. File doesn't exist")
		return "", err
	}

	return string(data), nil
}
