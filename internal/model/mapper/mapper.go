package mapper

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
)

type CruderMapper[TEntity any, TModel any] interface {
	ModelToResponse(entity *TEntity) *TModel
}

type Mapper[TEntity entity.Base, TModel model.BaseCollectionResponse] struct {
	Entity entity.Base
}

func (m *Mapper[TEntity, TModel]) ModelToResponse(entity *entity.Base) *TModel {
	return &TModel{ID: entity.ID}
}
