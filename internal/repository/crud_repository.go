package repository

import (
	"github.com/aikuci/go-subdivisions-id/internal/model"

	"gorm.io/gorm"
)

type CruderRepository[T any] interface {
	Find(db *gorm.DB) ([]T, error)
	FindById(db *gorm.DB, id int) ([]T, error)
	FindByIds(db *gorm.DB, ids []int) ([]T, error)
	FirstById(db *gorm.DB, id int) (*T, error)
}

type CrudRepository[T any, TId model.IdSingular, TIds model.IdPlural] struct {
	Repository[T, TId, TIds]
}

func NewCrudRepository[T any, TId model.IdSingular, TIds model.IdPlural]() *CrudRepository[T, TId, TIds] {
	return &CrudRepository[T, TId, TIds]{}
}
