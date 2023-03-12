package service

import "github.com/gofiber/fiber/v2"

type LandmarkService interface {
	SaveLandmark(c *fiber.Ctx) error
	UpdateLandmark(c *fiber.Ctx) error
	DeleteLandmark(c *fiber.Ctx) error
	GetLandmarkById(c *fiber.Ctx) error
	GetAllLandmarksByCityId(c *fiber.Ctx) error
}
