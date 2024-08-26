package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/aikuci/go-subdivisions-id/internal/model"
	apperror "github.com/aikuci/go-subdivisions-id/internal/pkg/error"
	"github.com/aikuci/go-subdivisions-id/internal/repository"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CruderUseCase[T any] interface {
	List(ctx context.Context, request model.ListRequest) ([]T, error)
	GetById(ctx context.Context, request model.GetByIDRequest[int]) ([]T, error)
	GetFirstById(ctx context.Context, request model.GetByIDRequest[int]) (*T, error)
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

func (uc *CrudUseCase[T]) List(ctx context.Context, request model.ListRequest) ([]T, error) {
	return wrapperPlural(
		newUseCase[T](ctx, uc.Log, uc.DB, request),
		func(ca *CallbackArgs[model.ListRequest]) ([]T, error) {
			return uc.Repository.Find(ca.tx)
		},
	)
}

func (uc *CrudUseCase[T]) GetById(ctx context.Context, request model.GetByIDRequest[int]) ([]T, error) {
	return wrapperPlural(
		newUseCase[T](ctx, uc.Log, uc.DB, request),
		func(ca *CallbackArgs[model.GetByIDRequest[int]]) ([]T, error) {
			return uc.Repository.FindById(ca.tx, ca.request.ID)
		},
	)
}

func (uc *CrudUseCase[T]) GetFirstById(ctx context.Context, request model.GetByIDRequest[int]) (*T, error) {
	return wrapperSingular(
		newUseCase[T](ctx, uc.Log, uc.DB, request),
		func(ca *CallbackArgs[model.GetByIDRequest[int]]) (*T, error) {
			db := ca.tx
			id := ca.request.ID
			collection, err := uc.Repository.FirstById(db, id)

			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					errorMessage := fmt.Sprintf("failed to get data with ID: %d", id)
					ca.log.Warn(err.Error(), zap.String("errorMessage", errorMessage))
					return nil, apperror.RecordNotFound(errorMessage)
				}

				ca.log.Warn(err.Error())
				return nil, err
			}

			return collection, nil
		},
	)
}
