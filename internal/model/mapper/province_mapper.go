package mapper

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	appmodel "github.com/aikuci/go-subdivisions-id/pkg/model"
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
	districtsMapper := NewDistrictMapper()
	districts := make([]model.DistrictResponse, len(province.Districts))
	for i, collection := range province.Districts {
		districts[i] = *districtsMapper.ModelToResponse(&collection)
	}

	return &model.ProvinceResponse{
		BaseCollectionResponse: appmodel.BaseCollectionResponse[int]{ID: province.ID},
		Code:                   province.Code,
		Name:                   province.Name,
		PostalCodes:            province.PostalCodes,
		Cities:                 cities,
		Districts:              districts,
	}
}
