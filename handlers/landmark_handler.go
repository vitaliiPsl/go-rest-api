package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vitaliiPsl/go-rest-api/service"
)

type LandmarkHandler struct {
	service service.LandmarkService
}

func NewAttractionsHandler(service service.LandmarkService) *LandmarkHandler {
	return &LandmarkHandler{service: service}
}

func (handler *LandmarkHandler) SaveLandmark(c *fiber.Ctx) error {
	return handler.service.SaveLandmark(c)
}

func (handler *LandmarkHandler) UpdateLandmark(c *fiber.Ctx) error {
	return handler.service.UpdateLandmark(c)
}

func (handler *LandmarkHandler) DeleteLandmark(c *fiber.Ctx) error {
	return handler.service.DeleteLandmark(c)
}

func (handler *LandmarkHandler) GetAllLandmarksByCityId(c *fiber.Ctx) error {
	return handler.service.GetAllLandmarksByCityId(c)
}

func (handler *LandmarkHandler) GetLandmarkById(c *fiber.Ctx) error {
	return handler.service.GetLandmarkById(c)
}

func RegisterLandmarksHandler(router fiber.Router, landmarkService service.LandmarkService) {
	// create new landmarks handler
	handler := NewAttractionsHandler(landmarkService)

	// register handlers to corresponding endpoints
	router.Post("cities/:cityId/landmarks", handler.SaveLandmark)

	router.Put("cities/:cityId/landmarks/:landmarkId", handler.UpdateLandmark)

	router.Delete("cities/:cityId/landmarks/:landmarkId", handler.DeleteLandmark)

	router.Get("cities/:cityId/landmarks", handler.GetAllLandmarksByCityId)

	router.Get("cities/:cityId/landmarks/:landmarkId", handler.GetLandmarkById)
}
