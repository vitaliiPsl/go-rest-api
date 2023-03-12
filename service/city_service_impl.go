package service

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/vitaliiPsl/go-rest-api/payload"
	"github.com/vitaliiPsl/go-rest-api/repository"
	"github.com/vitaliiPsl/go-rest-api/validation"
	"gorm.io/gorm"
)

type CitiesServiceImpl struct {
	validator      *validator.Validate
	cityRepository repository.CityRepository
}

func NewCityService(validator *validator.Validate, cityRepository repository.CityRepository) *CitiesServiceImpl {
	return &CitiesServiceImpl{validator: validator, cityRepository: cityRepository}
}

func (service *CitiesServiceImpl) SaveCity(c *fiber.Ctx) error {
	// parse body
	cityDto := payload.CityDto{}
	if err := c.BodyParser(&cityDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad request")
	}

	// validate dto
	if err := service.validator.Struct(cityDto); err != nil {
		errors := validation.CollectErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	city := payload.ToCityModel(&cityDto)
	if err := service.cityRepository.Save(city); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Something went wrong")
	}

	cityDto = *payload.ToCityDto(city)
	return c.Status(fiber.StatusCreated).JSON(cityDto)
}

func (service *CitiesServiceImpl) UpdateCity(c *fiber.Ctx) error {
	// parse id city id
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid path")
	}

	// check if city exist
	if _, err = service.cityRepository.FindById(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Couldn't find city with given id")
		}

		return fiber.NewError(fiber.StatusInternalServerError, "Somethign went wrong")
	}

	// parse body
	cityDto := payload.CityDto{}
	if err := c.BodyParser(&cityDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad request")
	}

	// validate dto
	if err := service.validator.Struct(cityDto); err != nil {
		errors := validation.CollectErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// replace city
	city := payload.ToCityModel(&cityDto)
	city.Id = id
	if err = service.cityRepository.Save(city); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Somethign went wrong")
	}

	cityDto = *payload.ToCityDto(city)
	return c.JSON(cityDto)
}

func (service *CitiesServiceImpl) DeleteCity(c *fiber.Ctx) error {
	// parse id city id
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid path")
	}

	// delete city
	if err := service.cityRepository.Delete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Somethign went wrong")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (service *CitiesServiceImpl) GetAllCities(c *fiber.Ctx) error {
	// get all cities
	cities, err := service.cityRepository.FindAll()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Something went wrong")
	}

	// map cities to dtos
	dtos := []payload.CityDto{}
	for _, city := range *cities {
		cityDto := payload.ToCityDto(&city)
		dtos = append(dtos, *cityDto)
	}

	return c.JSON(dtos)
}

func (service *CitiesServiceImpl) GetCityById(c *fiber.Ctx) error {
	// parse city id
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid path")
	}

	city, err := service.cityRepository.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Couldn't find city with given id")
		}

		return fiber.NewError(fiber.StatusInternalServerError, "Somethign went wrong")
	}

	return c.JSON(&city)
}
