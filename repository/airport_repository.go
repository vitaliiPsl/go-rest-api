package repository

import "github.com/vitaliiPsl/go-rest-api/model"

type AirportRepository interface {
	Save(landmark *model.Airport) error
	Delete(cityId, landmarkId int) error
	FindById(cityId, landmarkId int) (*model.Airport, error)
	FindAllByCityId(cityId int) (*[]model.Airport, error)
}