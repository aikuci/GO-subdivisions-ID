package usecase

import (
	"context"

	"github.com/aikuci/go-subdivisions-id/internal/model"
)

type CrudUseCaser[T any] interface {
	List(ctx context.Context, request *model.ListRequest) ([]T, error)
	GetByID(ctx context.Context, request *model.GetByIDRequest) (*T, error)
}
