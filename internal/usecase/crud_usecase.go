package usecase

import (
	"context"
	"errors"
	"fmt"

	appmodel "github.com/aikuci/go-subdivisions-id/pkg/model"
	apprepository "github.com/aikuci/go-subdivisions-id/pkg/repository"
	apperror "github.com/aikuci/go-subdivisions-id/pkg/util/error"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CruderUseCase[T any] interface {
	List(ctx context.Context, request appmodel.ListRequest) (*[]T, int64, error)
	GetById(ctx context.Context, request appmodel.GetByIDRequest[int]) (*[]T, int64, error)
	GetFirstById(ctx context.Context, request appmodel.GetByIDRequest[int]) (**T, int64, error)
}

type CrudUseCase[T any] struct {
	Log        *zap.Logger
	DB         *gorm.DB
	Repository apprepository.CruderRepository[T]
}

func NewCrudUseCase[T any](log *zap.Logger, db *gorm.DB, repository apprepository.CruderRepository[T]) *CrudUseCase[T] {
	return &CrudUseCase[T]{
		Log:        log,
		DB:         db,
		Repository: repository,
	}
}

func (uc *CrudUseCase[T]) List(ctx context.Context, request appmodel.ListRequest) (*[]T, int64, error) {
	callbackContext := &Context[appmodel.ListRequest, []T]{Data: *NewContextData(ctx, uc.Log, uc.DB, request)}
	return Wrapper[T](
		callbackContext,
		func(ctx *Context[appmodel.ListRequest, []T]) (ContextResult[[]T], error) {
			data := ctx.Data

			collections, total, err := uc.Repository.FindAndCount(data.DB)
			if err != nil {
				return ContextResult[[]T]{}, err
			}

			return ContextResult[[]T]{Collection: collections, Total: total}, nil
		},
	)
}

func (uc *CrudUseCase[T]) GetById(ctx context.Context, request appmodel.GetByIDRequest[int]) (*[]T, int64, error) {
	callbackContext := &Context[appmodel.GetByIDRequest[int], []T]{Data: *NewContextData(ctx, uc.Log, uc.DB, request)}
	return Wrapper[T](
		callbackContext,
		func(ctx *Context[appmodel.GetByIDRequest[int], []T]) (ContextResult[[]T], error) {
			data := ctx.Data

			collections, total, err := uc.Repository.FindAndCountById(data.DB, data.Request.ID)
			if err != nil {
				return ContextResult[[]T]{}, err
			}

			return ContextResult[[]T]{Collection: collections, Total: total}, nil
		},
	)
}

func (uc *CrudUseCase[T]) GetFirstById(ctx context.Context, request appmodel.GetByIDRequest[int]) (**T, int64, error) {
	callbackContext := &Context[appmodel.GetByIDRequest[int], *T]{Data: *NewContextData(ctx, uc.Log, uc.DB, request)}
	return Wrapper[T](
		callbackContext,
		func(ctx *Context[appmodel.GetByIDRequest[int], *T]) (ContextResult[*T], error) {
			data := ctx.Data

			id := data.Request.ID

			collection, err := uc.Repository.FirstById(data.DB, id)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					errorMessage := fmt.Sprintf("failed to get data with ID: %d", id)
					data.Log.Warn(err.Error(), zap.String("errorMessage", errorMessage))
					return ContextResult[*T]{}, apperror.RecordNotFound(errorMessage)
				}

				data.Log.Warn(err.Error())
				return ContextResult[*T]{}, err
			}

			return ContextResult[*T]{Collection: collection}, nil
		},
	)
}
