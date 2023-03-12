package main

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

// new JSON error handler
func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	// retrieve status code if it is a fiber error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	log.Printf("ERROR: %s\n", err.Error())
	return c.Status(code).JSON(fiber.Map{"message": err.Error()})
}

func main() {
	log.Println("Go rest api")

	app := fiber.New(fiber.Config{ErrorHandler: errorHandler})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Up and running",
		})
	})

	app.Listen(":3000")
}
