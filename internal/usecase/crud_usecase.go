package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/aikuci/go-subdivisions-id/internal/repository"
	appmodel "github.com/aikuci/go-subdivisions-id/pkg/model"
	apperror "github.com/aikuci/go-subdivisions-id/pkg/util/error"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CruderUseCase[T any] interface {
	List(ctx context.Context, request appmodel.ListRequest) ([]T, int64, error)
	GetById(ctx context.Context, request appmodel.GetByIDRequest[int]) ([]T, int64, error)
	GetFirstById(ctx context.Context, request appmodel.GetByIDRequest[int]) (*T, error)
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

func (uc *CrudUseCase[T]) List(ctx context.Context, request appmodel.ListRequest) ([]T, int64, error) {
	return wrapperPlural(
		newUseCase[T](ctx, uc.Log, uc.DB, request),
		func(ca *CallbackArgs[appmodel.ListRequest]) ([]T, int64, error) {
			return uc.Repository.FindAndCount(ca.tx)
		},
	)
}

func (uc *CrudUseCase[T]) GetById(ctx context.Context, request appmodel.GetByIDRequest[int]) ([]T, int64, error) {
	return wrapperPlural(
		newUseCase[T](ctx, uc.Log, uc.DB, request),
		func(ca *CallbackArgs[appmodel.GetByIDRequest[int]]) ([]T, int64, error) {
			return uc.Repository.FindAndCountById(ca.tx, ca.request.ID)
		},
	)
}

func (uc *CrudUseCase[T]) GetFirstById(ctx context.Context, request appmodel.GetByIDRequest[int]) (*T, error) {
	return wrapperSingular(
		newUseCase[T](ctx, uc.Log, uc.DB, request),
		func(ca *CallbackArgs[appmodel.GetByIDRequest[int]]) (*T, error) {
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
