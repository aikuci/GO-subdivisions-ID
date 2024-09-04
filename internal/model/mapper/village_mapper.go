package mapper

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	appmodel "github.com/aikuci/go-subdivisions-id/pkg/model"
)

type VillageMapper struct{}

func NewVillage() *VillageMapper {
	return &VillageMapper{}
}

func (m *VillageMapper) ModelToResponse(village *entity.Village) *model.VillageResponse {
	var province *model.ProvinceResponse
	if village.Province.ID > 0 {
		provinceMapper := NewProvince()
		province = provinceMapper.ModelToResponse(&village.Province)
	}
	var city *model.CityResponse
	if village.City.ID > 0 {
		cityMapper := NewCity()
		city = cityMapper.ModelToResponse(&village.City)
	}
	var district *model.DistrictResponse
	if village.District.ID > 0 {
		districtMapper := NewDistrict()
		district = districtMapper.ModelToResponse(&village.District)
	}

	return &model.VillageResponse{
		BaseCollectionResponse: appmodel.BaseCollectionResponse[int]{ID: village.ID},
		IDDistrict:             village.DistrictID,
		IDCity:                 village.CityID,
		IDProvince:             village.ProvinceID,
		Code:                   village.Code,
		Name:                   village.Name,
		PostalCodes:            village.PostalCodes,

		Province: province,
		City:     city,
		District: district,
	}
}
