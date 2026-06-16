package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"qr-service-go/configs"
)

func main() {
	cfg := configs.Load()

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "qr-service-go",
		})
	})

	log.Printf("QR service listening on port %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
