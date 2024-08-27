package repository

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	appmodel "github.com/aikuci/go-subdivisions-id/pkg/model"

	"gorm.io/gorm"
)

type DistrictRepository[TId appmodel.IdSingular, TIds appmodel.IdPlural] struct {
	Repository[entity.District, TId, TIds]
}

func NewDistrictRepository[TId appmodel.IdSingular, TIds appmodel.IdPlural]() *DistrictRepository[TId, TIds] {
	return &DistrictRepository[TId, TIds]{}
}

func (r *DistrictRepository[TId, TIds]) FirstByIdAndIdCityAndIdProvince(db *gorm.DB, id TId, id_city TId, id_province TId) (*entity.District, error) {
	return r.FirstBy(db, map[string]interface{}{"id": id, "id_city": id_city, "id_province": id_province})
}
