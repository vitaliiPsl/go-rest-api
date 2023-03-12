package payload

import "github.com/vitaliiPsl/go-rest-api/model"


type LandmarkDto struct {
	Id          int    `json:"id"`
	CityId      int    `json:"cityId"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Address     string `json:"address"`
}

func ToLandmarkModel(dto *LandmarkDto) *model.Landmark {
	return &model.Landmark{
		Id:          dto.Id,
		CityId:      dto.CityId,
		Name:        dto.Name,
		Description: dto.Description,
		Address:     dto.Address,
	}
}

func ToLandmarkDto(model *model.Landmark) *LandmarkDto {
	return &LandmarkDto{
		Id:          model.Id,
		CityId:      model.CityId,
		Name:        model.Name,
		Description: model.Description,
		Address:     model.Address,
	}
}
