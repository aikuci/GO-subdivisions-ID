package usecase

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/middleware/requestid"
	"github.com/aikuci/go-subdivisions-id/internal/pkg/slice"

	"github.com/gobeam/stringy"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UseCase[TEntity any, TRequest any] struct {
	Log     *zap.Logger
	DB      *gorm.DB
	Request TRequest
}

func newUseCase[TEntity any, TRequest any](log *zap.Logger, db *gorm.DB, request TRequest) *UseCase[TEntity, TRequest] {
	return &UseCase[TEntity, TRequest]{
		Log:     log,
		DB:      db,
		Request: request,
	}
}

type relations struct {
	pascal []string
	snake  []string
}

type CallbackParam[T any] struct {
	tx      *gorm.DB
	log     *zap.Logger
	request T
}

func wrapperSingular[TEntity any, TRequest any](ctx context.Context, uc *UseCase[TEntity, TRequest], fc func(cp *CallbackParam[TRequest]) (*TEntity, error)) (*TEntity, error) {
	log := uc.Log.With(zap.String("requestid", requestid.FromContext(ctx)))

	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	relations := getRelations[TEntity](uc.DB)
	tx, err := addRelations(tx, relations, uc.Request)
	if err != nil {
		return nil, err
	}

	collection, err := fc(&CallbackParam[TRequest]{tx: tx, log: log, request: uc.Request})
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Warn(err.Error(), zap.String("errorMessage", "failed to commit transaction"))
		return nil, fiber.ErrInternalServerError
	}

	return collection, nil
}

func wrapperPlural[TEntity any, TRequest any](ctx context.Context, uc *UseCase[TEntity, TRequest], fc func(cp *CallbackParam[TRequest]) ([]TEntity, error)) ([]TEntity, error) {
	log := uc.Log.With(zap.String("requestid", requestid.FromContext(ctx)))

	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	relations := getRelations[TEntity](uc.DB)
	log.Info(fmt.Sprintf("%v", relations))
	tx, err := addRelations(tx, relations, uc.Request)
	if err != nil {
		return nil, err
	}

	collections, err := fc(&CallbackParam[TRequest]{tx: tx, log: log, request: uc.Request})
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Warn(err.Error(), zap.String("errorMessage", "failed to commit transaction"))
		return nil, fiber.ErrInternalServerError
	}

	return collections, nil
}

func getRelations[TEntity any](db *gorm.DB) *relations {
	var collections []TEntity
	preloadDB := db.Session(&gorm.Session{
		Initialized:              true,
		DryRun:                   true,
		SkipHooks:                true,
		SkipDefaultTransaction:   true,
		DisableNestedTransaction: true,
	}).First(&collections)

	var relations_snake []string
	var relations_pascal []string
	for key := range preloadDB.Statement.Schema.Relationships.Relations {
		relations_pascal = append(relations_pascal, key)

		str := stringy.New(key)
		relations_snake = append(relations_snake, str.SnakeCase().ToLower())
	}

	return &relations{pascal: relations_pascal, snake: relations_snake}
}

func addRelations(db *gorm.DB, relations *relations, request any) (*gorm.DB, error) {
	v := reflect.ValueOf(request)
	if v.FieldByName("Include").IsValid() {
		if include, ok := v.FieldByName("Include").Interface().([]string); ok {
			for _, relation := range include {
				idx := slice.ArrayIndexOf(relations.snake, relation)
				if idx == -1 {
					return nil, fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Invalid relation '%v' provided. Available relation is '(%v)'.", relation, strings.Join(relations.snake, ", ")))
				}
				db = db.Preload(relations.pascal[idx])
			}
		}
	}
	return db, nil
}
