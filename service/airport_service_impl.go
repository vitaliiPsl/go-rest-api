package service

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/vitaliiPsl/go-rest-api/payload"
	"github.com/vitaliiPsl/go-rest-api/repository"
	"github.com/vitaliiPsl/go-rest-api/utils"
	"github.com/vitaliiPsl/go-rest-api/validation"
	"gorm.io/gorm"
)

type AirportServiceImpl struct {
	validator         *validator.Validate
	cityRepository    repository.CityRepository
	airportRepository repository.AirportRepository
}

func NewAirportService(
	validator *validator.Validate,
	cityRepository repository.CityRepository,
	airportRepository repository.AirportRepository,
) *AirportServiceImpl {
	return &AirportServiceImpl{validator: validator, cityRepository: cityRepository, airportRepository: airportRepository}
}

func (service *AirportServiceImpl) SaveAirport(c *fiber.Ctx) error {
	// get city id
	cityId, err := utils.GetPathId("cityId", c)
	if err != nil {
		return err
	}

	// find city
	if _, err = service.cityRepository.FindById(cityId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Couldn't find city")
		}

		return fiber.NewError(fiber.StatusInternalServerError, "Somethign went wrong")
	}

	// parse body
	airportDto := payload.AirportDto{}
	if err := c.BodyParser(&airportDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad request")
	}

	// validate dto
	if err := service.validator.Struct(airportDto); err != nil {
		errors := validation.CollectErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	airport := payload.ToAirportModel(&airportDto)
	airport.CityId = cityId
	if err := service.airportRepository.Save(airport); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Something went wrong")
	}

	airportDto = *payload.ToAirportDto(airport)
	return c.Status(fiber.StatusCreated).JSON(airportDto)
}

func (service *AirportServiceImpl) UpdateAirport(c *fiber.Ctx) error {
	// parse city and airport ids
	cityId, err := utils.GetPathId("cityId", c)
	if err != nil {
		return err
	}

	airportId, err := utils.GetPathId("airportId", c)
	if err != nil {
		return err
	}

	// check if airport exist
	if _, err = service.airportRepository.FindById(cityId, airportId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Couldn't find such airport")
		}

		return fiber.NewError(fiber.StatusInternalServerError, "Somethign went wrong")
	}

	// parse body
	airportDto := payload.AirportDto{}
	if err := c.BodyParser(&airportDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad request")
	}

	// validate dto
	if err := service.validator.Struct(airportDto); err != nil {
		errors := validation.CollectErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// replace airport
	airport := payload.ToAirportModel(&airportDto)
	airport.Id = airportId
	airport.CityId = cityId
	if err = service.airportRepository.Save(airport); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Somethign went wrong")
	}

	airportDto = *payload.ToAirportDto(airport)
	return c.JSON(airportDto)
}

func (service *AirportServiceImpl) DeleteAirport(c *fiber.Ctx) error {
	// parse city and airport ids
	cityId, err := utils.GetPathId("cityId", c)
	if err != nil {
		return err
	}

	airportId, err := utils.GetPathId("airportId", c)
	if err != nil {
		return err
	}

	// delete airport
	if err := service.airportRepository.Delete(cityId, airportId); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Somethign went wrong")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (service *AirportServiceImpl) GetAllAirportsByCityId(c *fiber.Ctx) error {
	// parse city id
	cityId, err := utils.GetPathId("cityId", c)
	if err != nil {
		return err
	}

	// get all airports
	airports, err := service.airportRepository.FindAllByCityId(cityId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Something went wrong")
	}

	// map airports to dtos
	dtos := []payload.AirportDto{}
	for _, airport := range *airports {
		airportDto := payload.ToAirportDto(&airport)
		dtos = append(dtos, *airportDto)
	}

	return c.JSON(dtos)
}

func (service *AirportServiceImpl) GetAirportById(c *fiber.Ctx) error {
	// parse city and airport ids
	cityId, err := utils.GetPathId("cityId", c)
	if err != nil {
		return err
	}

	airportId, err := utils.GetPathId("airportId", c)
	if err != nil {
		return err
	}

	airport, err := service.airportRepository.FindById(cityId, airportId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Couldn't find such airport")
		}

		return fiber.NewError(fiber.StatusInternalServerError, "Somethign went wrong")
	}

	return c.JSON(&airport)
}
