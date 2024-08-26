package usecase

import (
	"context"

	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/middleware/requestid"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UseCase[TEntity any, TRequest any] struct {
	Log     *zap.Logger
	DB      *gorm.DB
	Request TRequest
}

func NewUseCase[TEntity any, TRequest any](log *zap.Logger, db *gorm.DB, request TRequest) *UseCase[TEntity, TRequest] {
	return &UseCase[TEntity, TRequest]{
		Log:     log,
		DB:      db,
		Request: request,
	}
}

type CallbackParam[T any] struct {
	tx        *gorm.DB
	log       *zap.Logger
	relations []string
	request   T
}

func getRelations[TEntity any](db *gorm.DB) []string {
	var collections []TEntity
	preloadDB := db.Session(&gorm.Session{
		Initialized:              true,
		DryRun:                   true,
		SkipHooks:                true,
		SkipDefaultTransaction:   true,
		DisableNestedTransaction: true,
	}).First(&collections)
	var relations []string
	for key := range preloadDB.Statement.Schema.Relationships.Relations {
		relations = append(relations, key)
	}
	return relations
}

func WrapperSingular[TEntity any, TRequest any](ctx context.Context, uc *UseCase[TEntity, TRequest], callback func(cp *CallbackParam[TRequest]) (*TEntity, error)) (*TEntity, error) {
	log := uc.Log.With(zap.String("requestid", requestid.FromContext(ctx)))

	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	relations := getRelations[TEntity](uc.DB)

	collection, err := callback(&CallbackParam[TRequest]{tx: tx, log: log, relations: relations, request: uc.Request})
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Warn(err.Error(), zap.String("errorMessage", "failed to commit transaction"))
		return nil, fiber.ErrInternalServerError
	}

	return collection, nil
}

func WrapperPlural[TEntity any, TRequest any](ctx context.Context, uc *UseCase[TEntity, TRequest], callback func(cp *CallbackParam[TRequest]) ([]TEntity, error)) ([]TEntity, error) {
	log := uc.Log.With(zap.String("requestid", requestid.FromContext(ctx)))

	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	relations := getRelations[TEntity](uc.DB)

	collections, err := callback(&CallbackParam[TRequest]{tx: tx, log: log, relations: relations, request: uc.Request})
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Warn(err.Error(), zap.String("errorMessage", "failed to commit transaction"))
		return nil, fiber.ErrInternalServerError
	}

	return collections, nil
}
