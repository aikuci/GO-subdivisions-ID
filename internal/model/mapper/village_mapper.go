package mapper

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	appmodel "github.com/aikuci/go-subdivisions-id/pkg/model"
)

type VillageMapper struct{}

func NewVillageMapper() *VillageMapper {
	return &VillageMapper{}
}

func (m *VillageMapper) ModelToResponse(village *entity.Village) *model.VillageResponse {
	return &model.VillageResponse{
		BaseCollectionResponse: appmodel.BaseCollectionResponse[int]{ID: village.ID},
		IDDistrict:             village.DistrictID,
		IDCity:                 village.CityID,
		IDProvince:             village.ProvinceID,
		Code:                   village.Code,
		Name:                   village.Name,
		PostalCodes:            village.PostalCodes,
	}
}
