package main

import (
	"errors"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
	"github.com/vitaliiPsl/go-rest-api/database"
	"github.com/vitaliiPsl/go-rest-api/handlers"
	"github.com/vitaliiPsl/go-rest-api/repository"
	"github.com/vitaliiPsl/go-rest-api/service"
	"github.com/vitaliiPsl/go-rest-api/validation"
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

	// new fiber instance configured to use JSON error handler
	app := fiber.New(fiber.Config{ErrorHandler: errorHandler})

	// enable cors
	app.Use(cors.New())

	// register logger
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// set content type to JSON
	app.Use(func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		return c.Next()
	})

	// health endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Up and running",
		})
	})

	// create api router
	api := app.Group("/api/")

	// initialize repositories
	cityRepository := repository.NewCityRepository(database.Db)
	airportRepository := repository.NewAirportRepository(database.Db)
	landmarkRepository := repository.NewLandmarkRepository(database.Db)

	// initialize services
	cityService := service.NewCityService(validation.Validator, cityRepository)
	airportService := service.NewAirportService(validation.Validator, cityRepository, airportRepository)
	landmarkService := service.NewLandmarkService(validation.Validator, cityRepository, landmarkRepository)

	// register handlers
	handlers.RegisterCityHandlers(api, cityService)
	handlers.RegisterAirportsHandler(api, airportService)
	handlers.RegisterLandmarksHandler(api, landmarkService)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
