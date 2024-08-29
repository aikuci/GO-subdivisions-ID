package http

import (
	"context"
	"fmt"
	"math"
	"reflect"

	"github.com/aikuci/go-subdivisions-id/pkg/delivery/http/middleware/requestid"
	"github.com/aikuci/go-subdivisions-id/pkg/model"
	"github.com/aikuci/go-subdivisions-id/pkg/model/mapper"
	apperror "github.com/aikuci/go-subdivisions-id/pkg/util/error"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ControllerContextData struct {
	collection any
	total      int64
}

type ControllerContext[TRequest any, TEntity any, TModel any] struct {
	Log      *zap.Logger
	Ctx      context.Context
	FiberCtx *fiber.Ctx
	Request  TRequest
	Mapper   mapper.CruderMapper[TEntity, TModel]
	Data     ControllerContextData
}

func NewContext[TRequest any, TEntity any, TModel any](log *zap.Logger, fiberCtx *fiber.Ctx, mapper mapper.CruderMapper[TEntity, TModel]) *ControllerContext[TRequest, TEntity, TModel] {
	return &ControllerContext[TRequest, TEntity, TModel]{
		Log:      log,
		FiberCtx: fiberCtx,
		Mapper:   mapper,
	}
}

type Callback[TRequest any, TEntity any, TModel any] func(ctx *ControllerContext[TRequest, TEntity, TModel]) (any, int64, error)

func Wrapper[TRequest any, TEntity any, TModel any](ctx *ControllerContext[TRequest, TEntity, TModel], callback Callback[TRequest, TEntity, TModel]) error {
	ctx.Ctx = requestid.SetContext(ctx.FiberCtx.UserContext(), ctx.FiberCtx)
	ctx.Log = ctx.Log.With(zap.String("requestid", requestid.FromContext(ctx.Ctx)))

	requestParsed := new(TRequest)
	if err := parseRequest(ctx.FiberCtx, requestParsed); err != nil {
		return err
	}
	ctx.Request = *requestParsed

	collection, total, err := callback(ctx)
	if err != nil {
		ctx.Log.Warn(err.Error())
		return err
	}

	ctx.Data = ControllerContextData{collection: collection, total: total}
	return buildResponse(ctx)
}

func buildResponse[TRequest any, TEntity any, TModel any](ctx *ControllerContext[TRequest, TEntity, TModel]) error {
	data := ctx.Data

	collectionValue := reflect.ValueOf(data.collection).Elem()
	if collectionValue.Kind() == reflect.Slice {
		responses := make([]TModel, collectionValue.Len())
		for i := 0; i < collectionValue.Len(); i++ {
			item, ok := collectionValue.Index(i).Interface().(TEntity)
			if !ok {
				errorMessage := fmt.Sprintf("element at index %d in collection is not of type %T", i, (*TEntity)(nil))
				ctx.Log.Warn(errorMessage)
				return apperror.InternalServerError(errorMessage)
			}
			responses[i] = *ctx.Mapper.ModelToResponse(&item)
		}

		return ctx.FiberCtx.JSON(
			model.WebResponse[[]TModel]{
				Data: responses,
				Meta: &model.Meta{
					Page: generatePageMeta(ctx.Request, data.total),
				},
			},
		)
	}

	item, ok := data.collection.(*TEntity)
	if !ok {
		errorMessage := fmt.Sprintf("collection is not of type %T", (*TEntity)(nil))
		ctx.Log.Warn(errorMessage)
		return apperror.InternalServerError(errorMessage)
	}
	return ctx.FiberCtx.JSON(
		model.WebResponse[TModel]{
			Data: *ctx.Mapper.ModelToResponse(item),
		},
	)
}

func parseRequest(ctx *fiber.Ctx, request any) error {
	if err := ctx.QueryParser(request); err != nil {
		return err
	}

	if err := ctx.ParamsParser(request); err != nil {
		return err
	}

	method := ctx.Method()
	if method == "GET" {
		return nil
	}

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return nil
}

func generatePageMeta(request any, total int64) *model.PageMetadata {
	r := reflect.ValueOf(request)
	for i := 0; i < r.NumField(); i++ {
		if pagination, ok := r.Field(i).Interface().(model.PageRequest); ok {
			if pagination.Page > 0 && pagination.Size > 0 {
				return &model.PageMetadata{
					Page:      pagination.Page,
					Size:      pagination.Size,
					TotalItem: total,
					TotalPage: int64(math.Ceil(float64(total) / float64(pagination.Size))),
				}
			}
		}
	}
	return nil
}
