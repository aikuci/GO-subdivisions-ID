package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/middleware/requestid"
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
	UseCase UseCase[TEntity, TModel]

	Log        *zap.Logger
	DB         *gorm.DB
	Repository repository.CruderRepository[TEntity]
	Mapper     mapper.CruderMapper[TEntity, TModel]
}

func NewCrudUseCase[TEntity any, TModel any](log *zap.Logger, db *gorm.DB, repository repository.CruderRepository[TEntity], mapper mapper.CruderMapper[TEntity, TModel]) *CrudUseCase[TEntity, TModel] {
	useCase := NewUseCase(log, db, mapper) // BUG: Potential issue with state persistence
	// BUG: The NewUseCase function is expected to be called only once.
	// If NewCrudUseCase is invoked multiple times, it could lead to unexpected behavior due to the reuse of the UseCase instance from the initial call.
	// This might cause unintended side effects if the UseCase instance holds state or data that persists across multiple invocations.

	// NOTE: The UseCase instance is designed to be used from the initial call. Any modifications or data appending (e.g., via uc.UseCase.Log)
	// could lead to memory leaks or unintended data accumulation. Ensure that the UseCase instance is properly managed to avoid such issues.

	return &CrudUseCase[TEntity, TModel]{
		UseCase: *useCase,

		Log:        log,
		DB:         db,
		Repository: repository,
		Mapper:     mapper,
	}
}

func (uc *CrudUseCase[TEntity, TModel]) ListCorefunc(tx *gorm.DB) ([]TEntity, error) {
	collections, err := uc.Repository.Find(uc.UseCase.DB)

	uc.Log.Info("UseCase Core Fn") // BUG: Display Bug

	if err != nil {
		uc.UseCase.Log.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return collections, nil
}

func (uc *CrudUseCase[TEntity, TModel]) List(ctx context.Context, request *model.ListRequest) ([]TModel, error) {
	return uc.UseCase.WrapperPlural(
		ctx,
		uc.ListCorefunc,
	)
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
