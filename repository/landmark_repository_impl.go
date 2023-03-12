package repository

import (
	"github.com/vitaliiPsl/go-rest-api/model"
	"gorm.io/gorm"
)

type LandmarkRepositoryImpl struct {
	db *gorm.DB
}

func NewLandmarkRepository(db *gorm.DB) *LandmarkRepositoryImpl{
	return &LandmarkRepositoryImpl{db: db}
}

/*
	Save landmark and generate primary key if it doens't have one
*/
func (rep *LandmarkRepositoryImpl) Save(landmark *model.Landmark) error {
	err := rep.db.Save(landmark).Error

	return err
}

/*
	Delete landmark with given landmark and city ids
*/
func (rep *LandmarkRepositoryImpl) Delete(cityId, attractionId int) error {
	err := rep.db.Delete(&model.Landmark{}, "id = ? and city_id = ?", attractionId, cityId).Error

	return err
}

/*
	Get landmark with given landmark and city ids
*/
func (rep *LandmarkRepositoryImpl) FindById(cityId, attractionId int) (*model.Landmark, error) {
	var landmark model.Landmark

	if err := rep.db.First(&landmark, "id = ? and city_id = ?", attractionId, cityId).Error; err != nil {
		return nil, err
	}

	return &landmark, nil
}

/*
	Get all landmark by id of the city
*/
func (rep *LandmarkRepositoryImpl) FindAllByCityId(cityId int) (*[]model.Landmark, error) {
	var landmark []model.Landmark

	if err := rep.db.Find(&landmark, "city_id = ?", cityId).Error; err != nil {
		return nil, err
	}

	return &landmark, nil
}
