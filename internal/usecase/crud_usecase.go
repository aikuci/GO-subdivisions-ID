package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/pkg/slice"
	"github.com/aikuci/go-subdivisions-id/internal/repository"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CruderUseCase[T any, TRelation any] interface {
	List(ctx context.Context, request *model.ListRequest) ([]T, error)
	GetByID(ctx context.Context, request *model.GetByIDRequest[int]) ([]T, error)
	GetByIDs(ctx context.Context, request *model.GetByIDRequest[[]int]) ([]T, error)
	GetFirstByID(ctx context.Context, request *model.GetByIDRequest[int]) (*T, error)
}

type CrudUseCase[T any, TRelation []string] struct {
	Log        *zap.Logger
	DB         *gorm.DB
	Repository repository.CruderRepository[T]
	Relations  TRelation
}

func NewCrudUseCase[T any, TRelation []string](log *zap.Logger, db *gorm.DB, repository repository.CruderRepository[T], relations TRelation) *CrudUseCase[T, TRelation] {
	return &CrudUseCase[T, TRelation]{
		Log:        log,
		DB:         db,
		Repository: repository,
		Relations:  relations,
	}
}

func (uc *CrudUseCase[T, TRelation]) List(ctx context.Context, request *model.ListRequest) ([]T, error) {
	useCase := NewUseCase[T](uc.Log, uc.DB, request)

	return WrapperPlural(ctx, useCase, uc.listFn)
}
func (uc *CrudUseCase[T, TRelation]) listFn(cp *CallbackParam[*model.ListRequest]) ([]T, error) {
	db := cp.tx

	for _, include := range cp.request.Include {
		if slice.Contains(uc.Relations, include) {
			db = db.Preload(include)
		}
	}
	collections, err := uc.Repository.Find(db)

	if err != nil {
		cp.log.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return collections, nil
}

func (uc *CrudUseCase[T, TRelation]) GetByID(ctx context.Context, request *model.GetByIDRequest[int]) ([]T, error) {
	useCase := NewUseCase[T](uc.Log, uc.DB, request)

	return WrapperPlural(ctx, useCase, uc.getByIdFn)
}
func (uc *CrudUseCase[T, TRelation]) getByIdFn(cp *CallbackParam[*model.GetByIDRequest[int]]) ([]T, error) {
	db := cp.tx
	for _, include := range cp.request.Include {
		if slice.Contains(uc.Relations, include) {
			db = db.Preload(include)
		}
	}
	collections, err := uc.Repository.FindById(db, cp.request.ID)

	if err != nil {
		cp.log.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return collections, nil
}

func (uc *CrudUseCase[T, TRelation]) GetByIDs(ctx context.Context, request *model.GetByIDRequest[[]int]) ([]T, error) {
	useCase := NewUseCase[T](uc.Log, uc.DB, request)

	return WrapperPlural(ctx, useCase, uc.getByIdsFn)
}
func (uc *CrudUseCase[T, TRelation]) getByIdsFn(cp *CallbackParam[*model.GetByIDRequest[[]int]]) ([]T, error) {
	db := cp.tx
	for _, include := range cp.request.Include {
		if slice.Contains(uc.Relations, include) {
			db = db.Preload(include)
		}
	}
	collections, err := uc.Repository.FindByIds(db, cp.request.ID)

	if err != nil {
		cp.log.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return collections, nil
}

func (uc *CrudUseCase[T, TRelation]) GetFirstByID(ctx context.Context, request *model.GetByIDRequest[int]) (*T, error) {
	useCase := NewUseCase[T](uc.Log, uc.DB, request)

	return WrapperSingular(ctx, useCase, uc.getFirstByIdFn)
}
func (uc *CrudUseCase[T, TRelation]) getFirstByIdFn(cp *CallbackParam[*model.GetByIDRequest[int]]) (*T, error) {
	db := cp.tx
	for _, include := range cp.request.Include {
		if slice.Contains(uc.Relations, include) {
			db = db.Preload(include)
		}
	}
	id := cp.request.ID
	collection, err := uc.Repository.FirstById(db, id)

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
