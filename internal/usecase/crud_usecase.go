package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/repository"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CruderUseCase[T any] interface {
	List(ctx context.Context, request *model.ListRequest) ([]T, error)
	GetByID(ctx context.Context, request *model.GetByIDRequest[int]) ([]T, error)
	GetByIDs(ctx context.Context, request *model.GetByIDRequest[[]int]) ([]T, error)
	GetFirstByID(ctx context.Context, request *model.GetByIDRequest[int]) (*T, error)
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

func (uc *CrudUseCase[T]) List(ctx context.Context, request *model.ListRequest) ([]T, error) {
	useCase := NewUseCase[T](uc.Log, uc.DB, request)

	return WrapperPlural(ctx, useCase, uc.listFn)
}
func (uc *CrudUseCase[T]) listFn(cp *CallbackParam[*model.ListRequest]) ([]T, error) {
	collections, err := uc.Repository.Find(cp.tx)

	if err != nil {
		cp.log.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return collections, nil
}

func (uc *CrudUseCase[T]) GetByID(ctx context.Context, request *model.GetByIDRequest[int]) ([]T, error) {
	useCase := NewUseCase[T](uc.Log, uc.DB, request)

	return WrapperPlural(ctx, useCase, uc.getByIdFn)
}
func (uc *CrudUseCase[T]) getByIdFn(cp *CallbackParam[*model.GetByIDRequest[int]]) ([]T, error) {
	collections, err := uc.Repository.FindById(cp.tx, cp.request.ID)

	if err != nil {
		cp.log.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return collections, nil
}

func (uc *CrudUseCase[T]) GetByIDs(ctx context.Context, request *model.GetByIDRequest[[]int]) ([]T, error) {
	useCase := NewUseCase[T](uc.Log, uc.DB, request)

	return WrapperPlural(ctx, useCase, uc.getByIdsFn)
}
func (uc *CrudUseCase[T]) getByIdsFn(cp *CallbackParam[*model.GetByIDRequest[[]int]]) ([]T, error) {
	collections, err := uc.Repository.FindByIds(cp.tx, cp.request.ID)

	if err != nil {
		cp.log.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return collections, nil
}

func (uc *CrudUseCase[T]) GetFirstByID(ctx context.Context, request *model.GetByIDRequest[int]) (*T, error) {
	useCase := NewUseCase[T](uc.Log, uc.DB, request)

	return WrapperSingular(ctx, useCase, uc.getFirstByIdFn)
}
func (uc *CrudUseCase[T]) getFirstByIdFn(cp *CallbackParam[*model.GetByIDRequest[int]]) (*T, error) {
	id := cp.request.ID
	collection, err := uc.Repository.FirstById(cp.tx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cp.log.Warn(err.Error(), zap.String("errorMessage", fmt.Sprintf("failed to get data with ID: %d", id)))
			return nil, fiber.ErrNotFound
		}

		cp.log.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return collection, nil
}
