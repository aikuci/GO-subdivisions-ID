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

type CallbackParam struct {
	tx  *gorm.DB
	log *zap.Logger
}

func WrapperPlural[TEntity any, TModel any](ctx context.Context, uc *UseCase[TEntity, TModel], callback func(cp *CallbackParam) ([]TEntity, error)) ([]TModel, error) {
	log := uc.Log.With(zap.String("requestid", requestid.FromContext(ctx)))

	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	collections, err := callback(&CallbackParam{tx: tx, log: log})
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Warn(err.Error(), zap.String("errorMessage", "failed to commit transaction"))
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]TModel, len(collections))
	for i, collection := range collections {
		responses[i] = *uc.Mapper.ModelToResponse(&collection)
	}

	return responses, nil
}
