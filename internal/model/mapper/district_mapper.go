package mapper

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
)

type DistrictMapper struct{}

func NewDistrictMapper() *DistrictMapper {
	return &DistrictMapper{}
}

func (m *DistrictMapper) ModelToResponse(district *entity.District) *model.DistrictResponse {
	return &model.DistrictResponse{
		BaseCollectionResponse: model.BaseCollectionResponse[int]{ID: district.ID},
		IDCity:                 district.IDCity,
		IDProvince:             district.IDProvince,
		Code:                   district.Code,
		Name:                   district.Name,
		PostalCodes:            district.PostalCodes,
	}
}
