package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/model/mapper"
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

type CrudUseCase[TEntity any, TModel any] struct {
	Log        *zap.Logger
	DB         *gorm.DB
	Repository repository.CruderRepository[TEntity]
	Mapper     mapper.CruderMapper[TEntity, TModel]
}

func NewCrudUseCase[TEntity any, TModel any](log *zap.Logger, db *gorm.DB, repository repository.CruderRepository[TEntity], mapper mapper.CruderMapper[TEntity, TModel]) *CrudUseCase[TEntity, TModel] {
	return &CrudUseCase[TEntity, TModel]{
		Log:        log,
		DB:         db,
		Repository: repository,
		Mapper:     mapper,
	}
}

func (uc *CrudUseCase[TEntity, TModel]) List(ctx context.Context, request *model.ListRequest) ([]TModel, error) {
	useCase := NewUseCase(uc.Log, uc.DB, uc.Mapper, request)

	return WrapperPlural(
		ctx,
		useCase,
		uc.ListFn,
	)
}
func (uc *CrudUseCase[TEntity, TModel]) ListFn(cp *CallbackParam[*model.ListRequest]) ([]TEntity, error) {
	collections, err := uc.Repository.Find(cp.tx)

	if err != nil {
		cp.log.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return collections, nil
}

func (uc *CrudUseCase[TEntity, TModel]) GetByID(ctx context.Context, request *model.GetByIDRequest[int]) ([]TModel, error) {
	useCase := NewUseCase(uc.Log, uc.DB, uc.Mapper, request)

	return WrapperPlural(
		ctx,
		useCase,
		uc.GetByIdFn,
	)
}
func (uc *CrudUseCase[TEntity, TModel]) GetByIdFn(cp *CallbackParam[*model.GetByIDRequest[int]]) ([]TEntity, error) {
	collections, err := uc.Repository.FindById(cp.tx, cp.request.ID)

	if err != nil {
		cp.log.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return collections, nil
}

func (uc *CrudUseCase[TEntity, TModel]) GetByIDs(ctx context.Context, request *model.GetByIDRequest[[]int]) ([]TModel, error) {
	useCase := NewUseCase(uc.Log, uc.DB, uc.Mapper, request)

	return WrapperPlural(
		ctx,
		useCase,
		uc.GetByIdsFn,
	)
}
func (uc *CrudUseCase[TEntity, TModel]) GetByIdsFn(cp *CallbackParam[*model.GetByIDRequest[[]int]]) ([]TEntity, error) {
	collections, err := uc.Repository.FindByIds(cp.tx, cp.request.ID)

	if err != nil {
		cp.log.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return collections, nil
}

func (uc *CrudUseCase[TEntity, TModel]) GetFirstByID(ctx context.Context, request *model.GetByIDRequest[int]) (*TModel, error) {
	useCase := NewUseCase(uc.Log, uc.DB, uc.Mapper, request)

	return WrapperSingular(
		ctx,
		useCase,
		uc.GetFirstByIdFn,
	)
}
func (uc *CrudUseCase[TEntity, TModel]) GetFirstByIdFn(cp *CallbackParam[*model.GetByIDRequest[int]]) (*TEntity, error) {
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
