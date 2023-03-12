package repository

import "github.com/vitaliiPsl/go-rest-api/model"

type LandmarkRepository interface {
	Save(landmark *model.Landmark) error
	Delete(cityId, landmarkId int) error
	FindById(cityId, landmarkId int) (*model.Landmark, error)
	FindAllByCityId(cityId int) (*[]model.Landmark, error)
}
