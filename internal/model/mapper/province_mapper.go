package mapper

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
)

func ProvinceToResponse(province *entity.Province) *model.ProvinceResponse {
	return &model.ProvinceResponse{
		ID:          province.ID,
		Code:        province.Code,
		Name:        province.Name,
		PostalCodes: province.PostalCodes,
	}
}
