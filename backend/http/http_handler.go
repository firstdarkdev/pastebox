package http

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"haste/config"
	"haste/utils"
	"strings"
	"unicode/utf8"
)

// DocumentsRoute Main HTTP handler for documents
func DocumentsRoute(router fiber.Router, config config.RuntimeConfig) {

	// Post Endpoint to create a new document
	router.Post("/documents", func(ctx *fiber.Ctx) error {
		// Check if the file is a text file
		if !strings.HasPrefix(ctx.Get("Content-Type"), "text/") {
			log.Errorf("Invalid Content-Type:%s", ctx.Get("Content-Type"))
			return ctx.Status(fiber.StatusBadRequest).SendString("only text uploads allowed")
		}

		// Read the body
		content := ctx.Body()
		if len(content) == 0 {
			log.Error("Empty body")
			return ctx.Status(fiber.StatusBadRequest).SendString("empty body")
		}

		// Length check
		if len(content) > config.Conf.MaxLength {
			log.Error("Document too large")
			return ctx.Status(fiber.StatusBadRequest).SendString("Document too large")
		}

		// Fallback Binary check
		for _, b := range content {
			if b == 0 {
				log.Error("Tried to parse binary file")
				return ctx.Status(fiber.StatusBadRequest).SendString("binary data not allowed")
			}
		}

		// Check for UTF-8 validity
		if !utf8.Valid(content) {
			log.Error("Invalid UTF-8 text")
			return ctx.Status(fiber.StatusBadRequest).SendString("invalid UTF-8 text")
		}

		// Generate the paste key
		key := config.Generator.Generate(config.Conf.KeyLength)
		filename := utils.HashKey(key)

		// Store it
		_, err := config.Storage.Save(filename, content)
		if err != nil {
			log.Error("Failed to save file")
			return ctx.Status(fiber.StatusInternalServerError).SendString("failed to save file")
		}

		// Verbose Logging
		log.WithFields(log.Fields{
			"id":   key,
			"size": len(content),
		}).Debug("Saved document")

		// Return the key
		return ctx.JSON(fiber.Map{
			"key": key,
		})
	})

	// Raw Document endpoint
	router.Get("/raw/:key", func(ctx *fiber.Ctx) error {
		key := ctx.Params("key")
		filename := utils.HashKey(key)

		// Read the document from storage
		data, err := config.Storage.Read(filename)

		if err != nil {
			return ctx.Status(fiber.StatusNotFound).SendString("not found")
		}

		// Return the raw data
		return ctx.SendString(data)
	})

	// Read a document
	router.Get("/documents/:key", func(ctx *fiber.Ctx) error {
		key := ctx.Params("key")
		filename := utils.HashKey(key)

		// Read the document from storage
		data, err := config.Storage.Read(filename)
		if err != nil {
			return ctx.Status(fiber.StatusNotFound).SendString("not found")
		}

		// Verbose Logging
		log.WithFields(log.Fields{
			"id":   key,
			"data": data,
		}).Debug("Loaded Document")

		// Return the document
		return ctx.JSON(fiber.Map{
			"key":  key,
			"data": data,
		})
	})
}
