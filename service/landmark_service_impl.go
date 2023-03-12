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

type LandmarkServiceImpl struct {
	validator             *validator.Validate
	cityRepository        repository.CityRepository
	landmarkRepository repository.LandmarkRepository
}

func NewLandmarkService(
	validator *validator.Validate,
	cityRepository repository.CityRepository,
	landmarkRepository repository.LandmarkRepository,
) *LandmarkServiceImpl {
	return &LandmarkServiceImpl{validator: validator, cityRepository: cityRepository, landmarkRepository: landmarkRepository}
}

func (service *LandmarkServiceImpl) SaveLandmark(c *fiber.Ctx) error {
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
	landmarkDto := payload.LandmarkDto{}
	if err := c.BodyParser(&landmarkDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad request")
	}

	// validate dto
	if err := service.validator.Struct(landmarkDto); err != nil {
		errors := validation.CollectErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	landmark := payload.ToLandmarkModel(&landmarkDto)
	landmark.CityId = cityId
	if err := service.landmarkRepository.Save(landmark); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Something went wrong")
	}

	landmarkDto = *payload.ToLandmarkDto(landmark)
	return c.Status(fiber.StatusCreated).JSON(landmarkDto)
}

func (service *LandmarkServiceImpl) UpdateLandmark(c *fiber.Ctx) error {
	// parse city and landmark ids
	cityId, err := utils.GetPathId("cityId", c)
	if err != nil {
		return err
	}

	landmarkId, err := utils.GetPathId("landmarkId", c)
	if err != nil {
		return err
	}

	// check if landmark exist
	if _, err = service.landmarkRepository.FindById(cityId, landmarkId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Couldn't find such landmark")
		}

		return fiber.NewError(fiber.StatusInternalServerError, "Somethign went wrong")
	}

	// parse body
	landmarkDto := payload.LandmarkDto{}
	if err := c.BodyParser(&landmarkDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad request")
	}

	// validate dto
	if err := service.validator.Struct(landmarkDto); err != nil {
		errors := validation.CollectErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// replace landmark
	landmark := payload.ToLandmarkModel(&landmarkDto)
	landmark.Id = landmarkId
	landmark.CityId = cityId
	if err = service.landmarkRepository.Save(landmark); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Somethign went wrong")
	}

	landmarkDto = *payload.ToLandmarkDto(landmark)
	return c.JSON(landmarkDto)
}

func (service *LandmarkServiceImpl) DeleteLandmark(c *fiber.Ctx) error {
	// parse city and landmark ids
	cityId, err := utils.GetPathId("cityId", c)
	if err != nil {
		return err
	}

	landmarkId, err := utils.GetPathId("landmarkId", c)
	if err != nil {
		return err
	}

	// delete landmark
	if err := service.landmarkRepository.Delete(cityId, landmarkId); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Somethign went wrong")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (service *LandmarkServiceImpl) GetAllLandmarksByCityId(c *fiber.Ctx) error {
	// parse city id
	cityId, err := utils.GetPathId("cityId", c)
	if err != nil {
		return err
	}

	// get all landmarks
	landmarks, err := service.landmarkRepository.FindAllByCityId(cityId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Something went wrong")
	}

	// map landmarks to dtos
	dtos := []payload.LandmarkDto{}
	for _, landmark := range *landmarks {
		landmarkDto := payload.ToLandmarkDto(&landmark)
		dtos = append(dtos, *landmarkDto)
	}

	return c.JSON(dtos)
}

func (service *LandmarkServiceImpl) GetLandmarkById(c *fiber.Ctx) error {
	// parse city and landmark ids
	cityId, err := utils.GetPathId("cityId", c)
	if err != nil {
		return err
	}

	landmarkId, err := utils.GetPathId("landmarkId", c)
	if err != nil {
		return err
	}

	landmark, err := service.landmarkRepository.FindById(cityId, landmarkId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Couldn't find such landmark")
		}

		return fiber.NewError(fiber.StatusInternalServerError, "Somethign went wrong")
	}

	return c.JSON(&landmark)
}
