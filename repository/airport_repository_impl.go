package repository

import (
	"github.com/vitaliiPsl/go-rest-api/model"
	"gorm.io/gorm"
)

type AirportRepositoryImpl struct {
	db *gorm.DB
}

func NewAirportRepository(db *gorm.DB) *AirportRepositoryImpl{
	return &AirportRepositoryImpl{db: db}
}

/*
	Save airport and generate primary key if it doens't have one
*/
func (rep *AirportRepositoryImpl) Save(airport *model.Airport) error {
	err := rep.db.Save(airport).Error

	return err
}

/*
	Delete airport with given airport and city ids
*/
func (rep *AirportRepositoryImpl) Delete(cityId, airportId int) error {
	err := rep.db.Delete(&model.Airport{}, "id = ? and city_id = ?", airportId, cityId).Error

	return err
}

/*
	Get airport with given airport and city ids
*/
func (rep *AirportRepositoryImpl) FindById(cityId, airportId int) (*model.Airport, error) {
	var airport model.Airport

	if err := rep.db.First(&airport, "id = ? and city_id = ?", airportId, cityId).Error; err != nil {
		return nil, err
	}

	return &airport, nil
}

/*
	Get all airports by id of the city
*/
func (rep *AirportRepositoryImpl) FindAllByCityId(cityId int) (*[]model.Airport, error) {
	var airports []model.Airport

	if err := rep.db.Find(&airports, "city_id = ?", cityId).Error; err != nil {
		return nil, err
	}

	return &airports, nil
}
