package usecase

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/middleware/requestid"
	appmodel "github.com/aikuci/go-subdivisions-id/pkg/model"
	apperror "github.com/aikuci/go-subdivisions-id/pkg/util/error"
	"github.com/aikuci/go-subdivisions-id/pkg/util/slice"

	"github.com/gobeam/stringy"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Context[TEntity any, TRequest any] struct {
	Data   ContextData[TRequest]
	Result ContextResult[TEntity]
	Err    error
}

type ContextResult[T any] struct {
	Collection T
	Total      int64
}

type ContextData[T any] struct {
	Ctx     context.Context
	Log     *zap.Logger
	DB      *gorm.DB
	Request T
}

func NewContextData[TRequest any](ctx context.Context, log *zap.Logger, db *gorm.DB, request TRequest) *ContextData[TRequest] {
	return &ContextData[TRequest]{
		Ctx:     ctx,
		Log:     log,
		DB:      db,
		Request: request,
	}
}

type Callback[TEntity any, TRequest any] func(ctx *Context[TEntity, TRequest]) (ContextResult[TEntity], error)

func Wrapper[TEntity any, TRequest any](ctx *Context[TEntity, TRequest], callback Callback[TEntity, TRequest]) {
	ctx.Data.Log = ctx.Data.Log.With(zap.String("requestid", requestid.FromContext(ctx.Data.Ctx)))

	ctx.Data.DB = ctx.Data.DB.WithContext(ctx.Data.Ctx).Begin()
	defer ctx.Data.DB.Rollback()
	defer func() {
		if err := ctx.Data.DB.Commit().Error; err != nil {
			errorMessage := "failed to commit transaction"
			ctx.Data.Log.Warn(err.Error(), zap.String("errorMessage", errorMessage))
			ctx.Err = apperror.InternalServerError(errorMessage)
		}
	}()

	var err error
	ctx.Data.DB, err = addRelations(ctx.Data.Log, ctx.Data.DB, generateRelations[TEntity](ctx.Data.DB), ctx.Data.Request)
	if err != nil {
		ctx.Err = err
		return
	}
	ctx.Data.DB = addPagination(ctx.Data.DB, ctx.Data.Request)

	result, err := callback(ctx)
	ctx.Result = result
	ctx.Err = err
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

func addRelations(log *zap.Logger, db *gorm.DB, relations *relations, request any) (*gorm.DB, error) {
	r := reflect.ValueOf(request)
	if r.FieldByName("Include").IsValid() {
		if include, ok := r.FieldByName("Include").Interface().([]string); ok {
			for _, relation := range include {
				if strings.Contains(relation, ".") {
					// TODO: Check if the relation is valid
					str := stringy.New(relation)
					db = db.Preload(str.PascalCase().Delimited(".").Get())
				} else {
					idx := slice.ArrayIndexOf(relations.snake, relation)
					if idx == -1 {
						errorMessage := fmt.Sprintf("Invalid relation '%v' provided. Available relation is '(%v)'.", relation, strings.Join(relations.snake, ", "))
						log.Warn(errorMessage)
						return nil, apperror.BadRequest(errorMessage)
					}
					db = db.Preload(relations.pascal[idx])
				}
			}
		}
	}
	return db, nil
}

func addPagination(db *gorm.DB, request any) *gorm.DB {
	r := reflect.ValueOf(request)
	for i := 0; i < r.NumField(); i++ {
		if pagination, ok := r.Field(i).Interface().(appmodel.PageRequest); ok {
			if pagination.Page > 0 && pagination.Size > 0 {
				offset := (pagination.Page - 1) * pagination.Size
				return db.Offset(offset).Limit(pagination.Size)
			}
		}
	}
	return db
}
