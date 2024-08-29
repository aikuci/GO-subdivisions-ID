package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/aikuci/go-subdivisions-id/pkg/model"
	"github.com/aikuci/go-subdivisions-id/pkg/repository"
	apperror "github.com/aikuci/go-subdivisions-id/pkg/util/error"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CruderUseCase[T any] interface {
	List(ctx context.Context, request model.ListRequest) (*[]T, int64, error)
	GetById(ctx context.Context, request model.GetByIDRequest[int]) (*[]T, int64, error)
	GetFirstById(ctx context.Context, request model.GetByIDRequest[int]) (*T, int64, error)
}

type CrudUseCase[T any] struct {
	Log        *zap.Logger
	DB         *gorm.DB
	Repository repository.CruderRepository[T]
}

func NewCrudUseCase[T any](log *zap.Logger, db *gorm.DB, repository repository.CruderRepository[T]) *CrudUseCase[T] {
	return &CrudUseCase[T]{
		Log:        log,
		DB:         db,
		Repository: repository,
	}
}

func (uc *CrudUseCase[T]) List(ctx context.Context, request model.ListRequest) (*[]T, int64, error) {
	return Wrapper[T](
		NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *UseCaseContext[model.ListRequest]) (*[]T, int64, error) {
			collections, total, err := uc.Repository.FindAndCount(ctx.DB)
			return &collections, total, err
		},
	)
}

func (uc *CrudUseCase[T]) GetById(ctx context.Context, request model.GetByIDRequest[int]) (*[]T, int64, error) {
	return Wrapper[T](
		NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *UseCaseContext[model.GetByIDRequest[int]]) (*[]T, int64, error) {
			collections, total, err := uc.Repository.FindAndCountById(ctx.DB, ctx.Request.ID)
			return &collections, total, err
		},
	)
}

func (uc *CrudUseCase[T]) GetFirstById(ctx context.Context, request model.GetByIDRequest[int]) (*T, int64, error) {
	return Wrapper[T](
		NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *UseCaseContext[model.GetByIDRequest[int]]) (*T, int64, error) {
			id := ctx.Request.ID

			collection, err := uc.Repository.FirstById(ctx.DB, id)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					errorMessage := fmt.Sprintf("failed to get data with ID: %d", id)
					ctx.Log.Warn(err.Error(), zap.String("errorMessage", errorMessage))
					return nil, 0, apperror.RecordNotFound(errorMessage)
				}

				ctx.Log.Warn(err.Error())
				return nil, 0, err
			}

			return &collection, 1, nil
		},
	)
}
