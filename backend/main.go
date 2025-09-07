package main

import (
	log "github.com/sirupsen/logrus"
	"haste/config"
	"haste/http"
	"haste/keygen"
	"haste/storage"
)

var Config *config.Config
var RuntimeConfig *config.RuntimeConfig

func main() {
	// Load Config
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		// Fallback Haste Config
		cfg, err = config.LoadConfig("config.js")

		if err != nil {
			panic("No config.json or config.js found.")
		}
	}

	Config = cfg

	// Setup Logging
	logging := Config.Logging[0]

	if logging.Level == "verbose" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.SetFormatter(&log.TextFormatter{
		ForceColors:   logging.Colorize,
		FullTimestamp: true,
	})
	log.Printf("Loaded Config")

	RuntimeConfig = &config.RuntimeConfig{
		Conf: *cfg,
	}

	switch cfg.KeyGenerator.Type {
	case "random":
		RuntimeConfig.Generator = keygen.NewRandomKeyGenerator(Config.KeyGenerator.KeySpace)
		break

	case "phonetic":
		RuntimeConfig.Generator = &keygen.PhoneticKeyGenerator{}
		break

	case "dictionary":
		RuntimeConfig.Generator = &keygen.DictionaryKeyGenerator{}
		break
	}

	switch cfg.Storage.Type {
	case "file":
		RuntimeConfig.Storage = storage.NewFileStorage()
		break
	}

	http.StartHttpServer(*RuntimeConfig)
}
