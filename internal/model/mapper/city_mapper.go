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
	districtsMapper := NewDistrictMapper()
	districts := make([]model.DistrictResponse, len(city.Districts))
	for i, collection := range city.Districts {
		districts[i] = *districtsMapper.ModelToResponse(&collection)
	}

	return &model.CityResponse{
		BaseCollectionResponse: model.BaseCollectionResponse[int]{ID: city.ID},
		IDProvince:             city.ProvinceID,
		Code:                   city.Code,
		Name:                   city.Name,
		PostalCodes:            city.PostalCodes,
		Districts:              districts,
	}
}
