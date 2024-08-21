package mapper

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
)

type ProvinceMapper struct {
	Entity entity.Province
}

func NewProvinceMapper() *ProvinceMapper {
	return &ProvinceMapper{}
}

func (m *ProvinceMapper) ModelToResponse(province *entity.Province) *model.ProvinceResponse {
	return &model.ProvinceResponse{
		BaseCollectionResponse: model.BaseCollectionResponse{ID: province.ID},
		Code:                   province.Code,
		Name:                   province.Name,
		PostalCodes:            province.PostalCodes,
	}
}
