package repository

import (
	"github.com/vitaliiPsl/go-rest-api/model"
)

type CityRepository interface {
	Save(city *model.City) error
	Delete(id int) error
	FindById(id int) (*model.City, error)
	FindAll(id int) (*[]model.City, error)
}
