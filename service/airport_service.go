package service

import "github.com/gofiber/fiber/v2"

type AirportService interface {
	SaveAirport(c *fiber.Ctx) error
	UpdateAirport(c *fiber.Ctx) error
	DeleteAirport(c *fiber.Ctx) error
	GetAirportById(c *fiber.Ctx) error
	GetAllAirportsByCityId(c *fiber.Ctx) error
}
