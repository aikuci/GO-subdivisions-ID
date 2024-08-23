package usecase

import (
	"context"

	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/middleware/requestid"
	"github.com/aikuci/go-subdivisions-id/internal/model/mapper"
	"github.com/gofiber/fiber/v2"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UseCase[TEntity any, TModel any] struct {
	Log    *zap.Logger
	DB     *gorm.DB
	Mapper mapper.CruderMapper[TEntity, TModel]
}

func NewUseCase[TEntity any, TModel any](log *zap.Logger, db *gorm.DB, mapper mapper.CruderMapper[TEntity, TModel]) *UseCase[TEntity, TModel] {
	return &UseCase[TEntity, TModel]{
		Log:    log,
		DB:     db,
		Mapper: mapper,
	}
}

func (uc *UseCase[TEntity, TModel]) WrapperSingular(ctx context.Context, callback func(_ *gorm.DB) (*TEntity, error)) (*TEntity, error) {
	uc.Log = uc.Log.With(zap.String("requestid", requestid.FromContext(ctx)))

	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	result, err := callback(tx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.Warn(err.Error(), zap.String("errorMessage", "failed to commit transaction"))
		return nil, fiber.ErrInternalServerError
	}

	return result, nil
}

func (uc *UseCase[TEntity, TModel]) WrapperPlural(ctx context.Context, callback func(_ *gorm.DB) ([]TEntity, error)) ([]TModel, error) {
	uc.Log = uc.Log.With(zap.String("requestid", requestid.FromContext(ctx)))

	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	collections, err := callback(tx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.Warn(err.Error(), zap.String("errorMessage", "failed to commit transaction"))
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]TModel, len(collections))
	for i, collection := range collections {
		responses[i] = *uc.Mapper.ModelToResponse(&collection)
	}

	return responses, nil
}
