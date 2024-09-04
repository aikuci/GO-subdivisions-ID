package repository

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	appmodel "github.com/aikuci/go-subdivisions-id/pkg/model"
	apprepository "github.com/aikuci/go-subdivisions-id/pkg/repository"

	"gorm.io/gorm"
)

type Village[TId appmodel.IdSingular, TIds appmodel.IdPlural] struct {
	apprepository.Repository[entity.Village, TId, TIds]
}

func NewVillage[TId appmodel.IdSingular, TIds appmodel.IdPlural]() *Village[TId, TIds] {
	return &Village[TId, TIds]{}
}

func (r *Village[TId, TIds]) FirstByIdAndIdDistrictAndIdCityAndIdProvince(db *gorm.DB, id TId, id_district TId, id_city TId, id_province TId) (entity.Village, error) {
	return r.FirstBy(db, map[string]interface{}{"id": id, "id_district": id_district, "id_city": id_city, "id_province": id_province})
}
