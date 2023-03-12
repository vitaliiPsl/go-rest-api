package repository

import (
	"github.com/vitaliiPsl/go-rest-api/model"
	"gorm.io/gorm"
)


type CityRepositoryImpl struct {
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) *CityRepositoryImpl {
	return &CityRepositoryImpl{db: db}
}

/*
	Save city and generate primary key if it doens't have one
*/
func (rep *CityRepositoryImpl) Save(city *model.City) error {
	err := rep.db.Save(city).Error

	return err
}

/*
	Delete city with given id
*/
func (rep *CityRepositoryImpl) Delete(id int) error {
	err := rep.db.Delete(&model.City{}, id).Error

	return err
}

/*
	Get city with given id
*/
func (rep *CityRepositoryImpl) FindById(id int) (*model.City, error) {
	var city model.City

	if err := rep.db.First(&city, id).Error; err != nil {
		return nil, err
	}

	return &city, nil
}

/*
	Get all cities
*/
func (rep *CityRepositoryImpl) FindAll() (*[]model.City, error) {
	var cities []model.City

	if err := rep.db.Find(&cities).Error; err != nil {
		return nil, err
	}

	return &cities, nil
}
