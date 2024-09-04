package repository

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	appmodel "github.com/aikuci/go-subdivisions-id/pkg/model"
	apprepository "github.com/aikuci/go-subdivisions-id/pkg/repository"

	"gorm.io/gorm"
)

type District[TId appmodel.IdSingular, TIds appmodel.IdPlural] struct {
	apprepository.Repository[entity.District, TId, TIds]
}

func NewDistrict[TId appmodel.IdSingular, TIds appmodel.IdPlural]() *District[TId, TIds] {
	return &District[TId, TIds]{}
}

func (r *District[TId, TIds]) FirstByIdAndIdCityAndIdProvince(db *gorm.DB, id TId, id_city TId, id_province TId) (entity.District, error) {
	return r.FirstBy(db, map[string]interface{}{"id": id, "id_city": id_city, "id_province": id_province})
}
