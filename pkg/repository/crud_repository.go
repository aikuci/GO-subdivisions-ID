package repository

import (
	"github.com/aikuci/go-subdivisions-id/pkg/model"

	"gorm.io/gorm"
)

type CruderRepository[T any] interface {
	FirstById(db *gorm.DB, id int) (T, error)
	Find(db *gorm.DB) ([]T, error)
	FindById(db *gorm.DB, id int) ([]T, error)
	FindByIds(db *gorm.DB, ids []int) ([]T, error)
	FindAndCount(db *gorm.DB) ([]T, int64, error)
	FindAndCountById(db *gorm.DB, id int) ([]T, int64, error)
	FindAndCountByIds(db *gorm.DB, ids []int) ([]T, int64, error)
}

type Crud[T any, TId model.IdSingular, TIds model.IdPlural] struct {
	Repository[T, TId, TIds]
}

func NewCrud[T any, TId model.IdSingular, TIds model.IdPlural]() *Crud[T, TId, TIds] {
	return &Crud[T, TId, TIds]{}
}
