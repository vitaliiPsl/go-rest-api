package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vitaliiPsl/go-rest-api/service"
)

type AirportsHandler struct {
	service service.AirportService
}

func NewAirportsHandler(service service.AirportService) (*AirportsHandler) {
	return &AirportsHandler{service: service}
}

func (handler *AirportsHandler) SaveAirport(c *fiber.Ctx) error {
	return handler.service.SaveAirport(c)
}

func (handler *AirportsHandler) UpdateAirport(c *fiber.Ctx) error {
	return handler.service.UpdateAirport(c)
}

func (handler *AirportsHandler) DeleteAirport(c *fiber.Ctx) error {
	return handler.service.DeleteAirport(c)
}

func (handler *AirportsHandler) GetAllAirports(c *fiber.Ctx) error {
	return handler.service.GetAllAirportsByCityId(c)
}

func (handler *AirportsHandler) GetAirportById(c *fiber.Ctx) error {
	return handler.service.GetAirportById(c)
}

func RegisterAirportsHandler(router fiber.Router, airportsService service.AirportService){
	// create new attractions handler
	handler := NewAirportsHandler(airportsService)

	// register handlers to corresponding endpoints
	router.Post("cities/:cityId/airports", handler.SaveAirport)

	router.Put("cities/:cityId/airports/:airportId", handler.UpdateAirport)

	router.Delete("cities/:cityId/airports/:airportId", handler.DeleteAirport)

	router.Get("cities/:cityId/airports", handler.GetAllAirports)

	router.Get("cities/:cityId/airports/:airportId", handler.GetAirportById)
}