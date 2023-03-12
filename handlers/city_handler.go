package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vitaliiPsl/go-rest-api/service"
)

type CityHandler struct {
	service service.CityService
}

func NewCitiesHandler(service service.CityService) *CityHandler {
	return &CityHandler{service: service}
}

func (handler *CityHandler) SaveCity(c *fiber.Ctx) error {
	return handler.service.SaveCity(c)
}

func (handler *CityHandler) UpdateCity(c *fiber.Ctx) error {
	return handler.service.UpdateCity(c)
}

func (handler *CityHandler) DeleteCity(c *fiber.Ctx) error {
	return handler.service.DeleteCity(c)
}

func (handler *CityHandler) GetAllCities(c *fiber.Ctx) error {
	return handler.service.GetAllCities(c)
}

func (handler *CityHandler) GetCityById(c *fiber.Ctx) error {
	return handler.service.GetCityById(c)
}

func RegisterCityHandlers(router fiber.Router, cityService service.CityService) {
	// create new city handler
	handler := NewCitiesHandler(cityService)

	// register handlers to corresponding endpoints
	router.Post("cities", handler.SaveCity)

	router.Put("cities/:id", handler.UpdateCity)

	router.Delete("cities/:id", handler.DeleteCity)

	router.Get("cities", handler.GetAllCities)

	router.Get("cities/:id", handler.GetCityById)
}
