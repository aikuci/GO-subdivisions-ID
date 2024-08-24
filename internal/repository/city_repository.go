package repository

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"

	"gorm.io/gorm"
)

type CityRepository[TId model.IdSingular, TIds model.IdPlural] struct {
	Repository[entity.City, TId, TIds]
}

func NewCityRepository[TId model.IdSingular, TIds model.IdPlural]() *CityRepository[TId, TIds] {
	return &CityRepository[TId, TIds]{}
}

func (r *CityRepository[TId, TIds]) FindByIdProvince(db *gorm.DB, id_province TId) ([]entity.City, error) {
	return r.FindBy(db, map[string]interface{}{"id_province": id_province})
}

func (r *CityRepository[TId, TIds]) FindByIdProvinces(db *gorm.DB, id_province TIds) ([]entity.City, error) {
	return r.FindBy(db, map[string]interface{}{"id_province": id_province})
}

func (r *CityRepository[TId, TIds]) FirstByIdAndIdProvince(db *gorm.DB, id TId, id_province TId) (*entity.City, error) {
	return r.FirstBy(db, map[string]interface{}{"id": id, "id_province": id_province})
}
