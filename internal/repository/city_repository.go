package repository

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	appmodel "github.com/aikuci/go-subdivisions-id/pkg/model"
	apprepository "github.com/aikuci/go-subdivisions-id/pkg/repository"

	"gorm.io/gorm"
)

type City[TId appmodel.IdSingular, TIds appmodel.IdPlural] struct {
	apprepository.Repository[entity.City, TId, TIds]
}

func NewCity[TId appmodel.IdSingular, TIds appmodel.IdPlural]() *City[TId, TIds] {
	return &City[TId, TIds]{}
}

func (r *City[TId, TIds]) FirstByIdAndIdProvince(db *gorm.DB, id TId, id_province TId) (entity.City, error) {
	return r.FirstBy(db, map[string]interface{}{"id": id, "id_province": id_province})
}
