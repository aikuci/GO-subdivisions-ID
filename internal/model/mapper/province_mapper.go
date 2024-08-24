package mapper

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
)

type ProvinceMapper struct{}

func NewProvinceMapper() *ProvinceMapper {
	return &ProvinceMapper{}
}

func (m *ProvinceMapper) ModelToResponse(province *entity.Province) *model.ProvinceResponse {
	citiesMapper := NewCityMapper()
	cities := make([]model.CityResponse, len(province.Cities))
	for i, collection := range province.Cities {
		cities[i] = *citiesMapper.ModelToResponse(&collection)
	}

	return &model.ProvinceResponse{
		BaseCollectionResponse: model.BaseCollectionResponse[int]{ID: province.ID},
		Code:                   province.Code,
		Name:                   province.Name,
		PostalCodes:            province.PostalCodes,
		Cities:                 cities,
	}
}
