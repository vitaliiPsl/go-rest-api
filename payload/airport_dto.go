package payload

import "github.com/vitaliiPsl/go-rest-api/model"

type AirportDto struct {
	Id      int    `json:"id"`
	CityId  int    `json:"cityId"`
	Name    string `json:"name" validate:"required"`
	Code    string `json:"code" validate:"required"`
	Opened  int    `json:"opened"`
	Address string `json:"address"`
}

func ToAirportModel(dto *AirportDto) *model.Airport {
	return &model.Airport{
		Id:      dto.Id,
		CityId:  dto.CityId,
		Name:    dto.Name,
		Code:    dto.Code,
		Opened:  dto.Opened,
		Address: dto.Address,
	}
}

func ToAirportDto(airport *model.Airport) *AirportDto {
	return &AirportDto{
		Id:      airport.Id,
		CityId:  airport.CityId,
		Name:    airport.Name,
		Code:    airport.Code,
		Opened:  airport.Opened,
		Address: airport.Address,
	}
}
