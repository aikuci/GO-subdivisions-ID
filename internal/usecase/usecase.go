package usecase

import (
	"context"

	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/middleware/requestid"
	"github.com/aikuci/go-subdivisions-id/internal/model/mapper"
	"github.com/gofiber/fiber/v2"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UseCase[TEntity any, TModel any, TRequest any] struct {
	Log     *zap.Logger
	DB      *gorm.DB
	Mapper  mapper.CruderMapper[TEntity, TModel]
	Request TRequest
}

func NewUseCase[TEntity any, TModel any, TRequest any](log *zap.Logger, db *gorm.DB, mapper mapper.CruderMapper[TEntity, TModel], request TRequest) *UseCase[TEntity, TModel, TRequest] {
	return &UseCase[TEntity, TModel, TRequest]{
		Log:     log,
		DB:      db,
		Mapper:  mapper,
		Request: request,
	}
}

type CallbackParam[TRequest any] struct {
	tx      *gorm.DB
	log     *zap.Logger
	request TRequest
}

func WrapperSingular[TEntity any, TModel any, TRequest any](ctx context.Context, uc *UseCase[TEntity, TModel, TRequest], callback func(cp *CallbackParam[TRequest]) (*TEntity, error)) (*TModel, error) {
	log := uc.Log.With(zap.String("requestid", requestid.FromContext(ctx)))

	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	collection, err := callback(&CallbackParam[TRequest]{tx: tx, log: log, request: uc.Request})
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Warn(err.Error(), zap.String("errorMessage", "failed to commit transaction"))
		return nil, fiber.ErrInternalServerError
	}

	return uc.Mapper.ModelToResponse(collection), nil
}

func WrapperPlural[TEntity any, TModel any, TRequest any](ctx context.Context, uc *UseCase[TEntity, TModel, TRequest], callback func(cp *CallbackParam[TRequest]) ([]TEntity, error)) ([]TModel, error) {
	log := uc.Log.With(zap.String("requestid", requestid.FromContext(ctx)))

	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	collections, err := callback(&CallbackParam[TRequest]{tx: tx, log: log, request: uc.Request})
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
