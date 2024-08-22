package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/middleware/requestid"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/model/mapper"
	"github.com/aikuci/go-subdivisions-id/internal/repository"

	"github.com/go-playground/validator/v10"
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
	Validate   *validator.Validate
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
	logger := uc.Log.With(zap.String(string("requestid"), requestid.FromContext(ctx)))

	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	collections, err := uc.Repository.Find(tx)
	if err != nil {
		logger.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		logger.Warn(err.Error(), zap.String("errorMessage", "failed to commit transaction"))
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]TModel, len(collections))
	for i, collection := range collections {
		responses[i] = *uc.Mapper.ModelToResponse(&collection)
	}

	return responses, nil
}

func (uc *CrudUseCase[TEntity, TModel]) GetByID(ctx context.Context, request *model.GetByIDRequest[int]) ([]TModel, error) {
	logger := uc.Log.With(zap.String(string("requestid"), requestid.FromContext(ctx)))

	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	collections, err := uc.Repository.FindById(tx, request.ID)
	if err != nil {
		logger.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		logger.Warn(err.Error(), zap.String("errorMessage", "failed to commit transaction"))
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]TModel, len(collections))
	for i, collection := range collections {
		responses[i] = *uc.Mapper.ModelToResponse(&collection)
	}

	return responses, nil
}

func (uc *CrudUseCase[TEntity, TModel]) GetByIDs(ctx context.Context, request *model.GetByIDRequest[[]int]) ([]TModel, error) {
	logger := uc.Log.With(zap.String(string("requestid"), requestid.FromContext(ctx)))

	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	collections, err := uc.Repository.FindByIds(tx, request.ID)
	if err != nil {
		logger.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		logger.Warn(err.Error(), zap.String("errorMessage", "failed to commit transaction"))
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]TModel, len(collections))
	for i, collection := range collections {
		responses[i] = *uc.Mapper.ModelToResponse(&collection)
	}

	return responses, nil
}

func (uc *CrudUseCase[TEntity, TModel]) GetFirstByID(ctx context.Context, request *model.GetByIDRequest[int]) (*TModel, error) {
	logger := uc.Log.With(zap.String(string("requestid"), requestid.FromContext(ctx)))

	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	id := request.ID
	collection, err := uc.Repository.FirstById(tx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warn(err.Error(), zap.String("errorMessage", fmt.Sprintf("failed to get data with ID: %d", id)))
			return nil, fiber.ErrNotFound
		}

		logger.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		logger.Warn(err.Error(), zap.String("errorMessage", "failed to commit transaction"))
		return nil, fiber.ErrInternalServerError
	}

	return uc.Mapper.ModelToResponse(collection), nil
}
