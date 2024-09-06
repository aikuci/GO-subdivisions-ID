package usecase

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/aikuci/go-subdivisions-id/pkg/model"
	apperror "github.com/aikuci/go-subdivisions-id/pkg/util/error"
	applog "github.com/aikuci/go-subdivisions-id/pkg/util/log"
	"github.com/aikuci/go-subdivisions-id/pkg/util/slice"

	"github.com/gobeam/stringy"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Context[T any] struct {
	Ctx     context.Context
	Log     *zap.Logger
	DB      *gorm.DB
	Request T
}

func NewContext[T any](ctx context.Context, log *zap.Logger, db *gorm.DB, request T) *Context[T] {
	return &Context[T]{
		Ctx:     ctx,
		Log:     log,
		DB:      db,
		Request: request,
	}
}

type Callback[TEntity any, TRequest any, TResult any] func(ctx *Context[TRequest]) (*TResult, int64, error)

func Wrapper[TEntity any, TRequest any, TResult any](ctx *Context[TRequest], callback Callback[TEntity, TRequest, TResult]) (*TResult, int64, error) {
	var apperr *apperror.CustomErrorResponse
	ctx.DB, apperr = addRelations(ctx.DB, generateRelations[TEntity](ctx.DB), ctx.Request)
	if apperr != nil {
		if apperr.HTTPCode != 500 {
			return nil, 0, apperr
		}

		errorMessage := "failed to process its relation"
		applog.Write(ctx.Log, ctx.Ctx, fmt.Sprintf("%v: ", errorMessage), apperr)
		return nil, 0, apperror.InternalServerError(errorMessage)
	}
	ctx.DB = addPagination(ctx.DB, ctx.Request)

	ctx.DB = ctx.DB.WithContext(ctx.Ctx).Begin()
	defer ctx.DB.Rollback()

	collections, total, err := callback(ctx)
	if err != nil {
		if apperr, ok := err.(*apperror.CustomErrorResponse); ok {
			return nil, 0, apperr
		}

		errorMessage := "failed to process"
		applog.Write(ctx.Log, ctx.Ctx, fmt.Sprintf("%v: ", errorMessage), err)
		return nil, 0, apperror.InternalServerError(errorMessage)
	}

	if err := ctx.DB.Commit().Error; err != nil {
		errorMessage := "failed to commit transaction"
		applog.Write(ctx.Log, ctx.Ctx, fmt.Sprintf("%v: ", errorMessage), err)
		return nil, 0, apperror.InternalServerError(errorMessage)
	}

	return collections, total, nil
}

type relations struct {
	pascal []string
	snake  []string
}

// generateRelations uses generics to collect relationships from a database model.
func generateRelations[TEntity any](db *gorm.DB) *relations {
	var collection TEntity
	return collectRelations(db, collection)
}

// collectRelations extracts relationship information from a model's schema.
func collectRelations(db *gorm.DB, collection any) *relations {
	preloadDB := db.Session(&gorm.Session{
		Initialized:              true,
		DryRun:                   true,
		SkipHooks:                true,
		SkipDefaultTransaction:   true,
		DisableNestedTransaction: true,
	}).First(collection)

	var relations_snake []string
	var relations_pascal []string
	for key := range preloadDB.Statement.Schema.Relationships.Relations {
		if strings.HasPrefix(key, "_") {
			continue
		}

		relations_pascal = append(relations_pascal, key)

		str := stringy.New(key)
		relations_snake = append(relations_snake, str.SnakeCase().ToLower())
	}

	return &relations{pascal: relations_pascal, snake: relations_snake}
}

func addRelations(db *gorm.DB, relations *relations, request any) (*gorm.DB, *apperror.CustomErrorResponse) {
	r := reflect.ValueOf(request)

	if !r.FieldByName("Include").IsValid() {
		return db, nil
	}

	include, ok := r.FieldByName("Include").Interface().([]string)
	if !ok {
		return nil, apperror.InternalServerError("Invalid type for 'Include' field. Expected []string.")
	}

	for _, relation := range include {
		if strings.Contains(relation, ".") {
			// TODO: Implement relation validation logic, Validate if relPath is a valid relation
			str := stringy.New(relation)
			db = db.Preload(str.PascalCase().Delimited(".").Get())
		} else {
			idx := slice.ArrayIndexOf(relations.snake, relation)
			if idx == -1 {
				return nil, apperror.BadRequest(fmt.Sprintf("Invalid relation: %v provided. Available relations are [%v].", relation, strings.Join(relations.snake, ", ")))
			}
			db = db.Preload(relations.pascal[idx])
		}
	}

	return db, nil
}

func addPagination(db *gorm.DB, request any) *gorm.DB {
	r := reflect.ValueOf(request)

	pageRequestField := r.FieldByName("PageRequest")
	if !pageRequestField.IsValid() || pageRequestField.Kind() != reflect.Struct {
		return db
	}

	pagination, ok := pageRequestField.Interface().(model.PageRequest)
	if !ok {
		return db
	}

	if pagination.Page > 0 && pagination.Size > 0 {
		offset := (pagination.Page - 1) * pagination.Size
		return db.Offset(offset).Limit(pagination.Size)
	}

	return db
}
