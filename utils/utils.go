package utils

import "github.com/gofiber/fiber/v2"

func GetPathId(pathVariable string, c *fiber.Ctx) (int, error) {
	id, err := c.ParamsInt(pathVariable)
	if err != nil {
		return 0, fiber.NewError(fiber.StatusBadRequest, "Invalid path")
	}

	return id, nil
}