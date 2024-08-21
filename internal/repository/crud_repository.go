package repository

import "go.uber.org/zap"

type CrudRepository[T any] struct {
	Repository[T]
	Log *zap.Logger
}

func NewCrudRepository[T any](log *zap.Logger) *CrudRepository[T] {
	return &CrudRepository[T]{
		Log: log,
	}
}
