package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	log "github.com/sirupsen/logrus"
	"haste/config"
)

// StartHttpServer Start the HTTP Server
func StartHttpServer(runtimeConfig config.RuntimeConfig) {
	// Setup HTTP Server
	log.Debug("Initializing HTTP Server")
	app := fiber.New()

	// Allow use from any domain
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// Logging and routes
	app.Use(logger.New())
	DocumentsRoute(app, runtimeConfig)

	app.Static("/", "./web/dist")

	app.Use(func(c *fiber.Ctx) error {
		return c.SendFile("./web/dist/index.html")
	})

	// Start the server
	addr := fmt.Sprintf("%s:%d", runtimeConfig.Conf.Host, runtimeConfig.Conf.Port)
	log.Infof("Server running on %s\n", addr)
	err := app.Listen(addr)

	if err != nil {
		return
	}
}
