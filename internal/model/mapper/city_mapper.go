package mapper

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	appmodel "github.com/aikuci/go-subdivisions-id/pkg/model"
)

type CityMapper struct{}

func NewCity() *CityMapper {
	return &CityMapper{}
}

func (m *CityMapper) ModelToResponse(city *entity.City) *model.CityResponse {
	var province *model.ProvinceResponse
	if city.Province.ID > 0 {
		provinceMapper := NewProvince()
		province = provinceMapper.ModelToResponse(&city.Province)
	}

	districtsMapper := NewDistrict()
	districts := make([]model.DistrictResponse, len(city.Districts))
	for i, collection := range city.Districts {
		districts[i] = *districtsMapper.ModelToResponse(&collection)
	}
	villagesMapper := NewVillage()
	villages := make([]model.VillageResponse, len(city.Villages))
	for i, collection := range city.Villages {
		villages[i] = *villagesMapper.ModelToResponse(&collection)
	}

	return &model.CityResponse{
		BaseCollectionResponse: appmodel.BaseCollectionResponse[int]{ID: city.ID},
		IDProvince:             city.ProvinceID,
		Code:                   city.Code,
		Name:                   city.Name,
		PostalCodes:            city.PostalCodes,

		Province:  province,
		Districts: districts,
		Villages:  villages,
	}
}
