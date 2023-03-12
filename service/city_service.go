package service

import "github.com/gofiber/fiber/v2"

type CityService interface {
	SaveCity(c *fiber.Ctx) error
	UpdateCity(c *fiber.Ctx) error
	DeleteCity(c *fiber.Ctx) error
	GetAllCities(c *fiber.Ctx) error
	GetCityById(c *fiber.Ctx) error
}
