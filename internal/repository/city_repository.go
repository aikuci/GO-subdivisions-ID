package repository

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	appmodel "github.com/aikuci/go-subdivisions-id/pkg/model"

	"gorm.io/gorm"
)

type CityRepository[TId appmodel.IdSingular, TIds appmodel.IdPlural] struct {
	Repository[entity.City, TId, TIds]
}

func NewCityRepository[TId appmodel.IdSingular, TIds appmodel.IdPlural]() *CityRepository[TId, TIds] {
	return &CityRepository[TId, TIds]{}
}

func (r *CityRepository[TId, TIds]) FirstByIdAndIdProvince(db *gorm.DB, id TId, id_province TId) (*entity.City, error) {
	return r.FirstBy(db, map[string]interface{}{"id": id, "id_province": id_province})
}
