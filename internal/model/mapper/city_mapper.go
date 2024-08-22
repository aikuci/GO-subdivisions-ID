package mapper

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
)

type CityMapper struct{}

func NewCityMapper() *CityMapper {
	return &CityMapper{}
}

func (m *CityMapper) ModelToResponse(city *entity.City) *model.CityResponse {
	return &model.CityResponse{
		BaseCollectionResponse: model.BaseCollectionResponse[int]{ID: city.ID},
		IDProvince:             city.IDProvince,
		Code:                   city.Code,
		Name:                   city.Name,
		PostalCodes:            city.PostalCodes,
	}
}
