package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CruderRepository[T any] interface {
	FindById(db *gorm.DB, entity *T, id any) error
	FindAll(db *gorm.DB) ([]T, error)
}

type CrudRepository[T any] struct {
	Repository[T]
	Log *zap.Logger
}

func NewCrudRepository[T any](log *zap.Logger) *CrudRepository[T] {
	return &CrudRepository[T]{
		Log: log,
	}
}
