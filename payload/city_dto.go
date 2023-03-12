package payload

import "github.com/vitaliiPsl/go-rest-api/model"


type CityDto struct {
	Id      int    `json:"id"`
	Name    string  `json:"name" validate:"required"`
	Country string  `json:"country" validate:"required"`
	Region  string  `json:"region"`
	Geolat  float32 `json:"latitude"`
	Geolng  float32 `json:"longitude"`
}

func ToCityModel(dto *CityDto) *model.City {
	return &model.City{
		Id:      dto.Id,
		Name:    dto.Name,
		Country: dto.Country,
		Region:  dto.Region,
		Geolat:  dto.Geolat,
		Geolng:  dto.Geolng,
	}
}

func ToCityDto(model *model.City) *CityDto {
	return &CityDto{
		Id:      model.Id,
		Name:    model.Name,
		Country: model.Country,
		Region:  model.Region,
		Geolat:  model.Geolat,
		Geolng:  model.Geolng,
	}
}