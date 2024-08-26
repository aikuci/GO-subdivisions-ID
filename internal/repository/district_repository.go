package repository

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"

	"gorm.io/gorm"
)

type DistrictRepository[TId model.IdSingular, TIds model.IdPlural] struct {
	Repository[entity.District, TId, TIds]
}

func NewDistrictRepository[TId model.IdSingular, TIds model.IdPlural]() *DistrictRepository[TId, TIds] {
	return &DistrictRepository[TId, TIds]{}
}

func (r *DistrictRepository[TId, TIds]) FirstByIdAndIdCityAndIdProvince(db *gorm.DB, id TId, id_city TId, id_province TId) (*entity.District, error) {
	return r.FirstBy(db, map[string]interface{}{"id": id, "id_city": id_city, "id_province": id_province})
}
