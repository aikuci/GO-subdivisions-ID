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
	return &model.ProvinceResponse{
		BaseCollectionResponse: model.BaseCollectionResponse[int]{ID: province.ID},
		Code:                   province.Code,
		Name:                   province.Name,
		PostalCodes:            province.PostalCodes,
	}
}
