package handler

import (
	"context"
	"math"
	"reflect"

	"github.com/aikuci/go-subdivisions-id/pkg/model"
	"github.com/aikuci/go-subdivisions-id/pkg/model/mapper"
	"github.com/aikuci/go-subdivisions-id/pkg/util/context/requestid"

	"github.com/gofiber/fiber/v2"
)

type ContextData struct {
	collection any
	total      int64
}

type Context[TRequest any, TEntity any, TModel any] struct {
	Ctx      context.Context
	FiberCtx *fiber.Ctx
	Request  TRequest
	Mapper   mapper.CruderMapper[TEntity, TModel]
	Data     ContextData
}

func NewContext[TRequest any, TEntity any, TModel any](fiberCtx *fiber.Ctx, mapper mapper.CruderMapper[TEntity, TModel]) *Context[TRequest, TEntity, TModel] {
	return &Context[TRequest, TEntity, TModel]{
		FiberCtx: fiberCtx,
		Mapper:   mapper,
	}
}

type Callback[TRequest any, TEntity any, TModel any] func(ctx *Context[TRequest, TEntity, TModel]) (any, int64, error)

func Wrapper[TRequest any, TEntity any, TModel any](ctx *Context[TRequest, TEntity, TModel], callback Callback[TRequest, TEntity, TModel]) error {
	ctx.Ctx = requestid.SetContext(ctx.FiberCtx.UserContext(), ctx.FiberCtx)
	ctx.FiberCtx.SetUserContext(ctx.Ctx)

	requestParsed := new(TRequest)
	if err := parseRequest(ctx.FiberCtx, requestParsed); err != nil {
		return err
	}
	ctx.Request = *requestParsed

	collection, total, err := callback(ctx)
	if err != nil {
		return err
	}

	ctx.Data = ContextData{collection: collection, total: total}
	return buildResponse(ctx)
}

func buildResponse[TRequest any, TEntity any, TModel any](ctx *Context[TRequest, TEntity, TModel]) error {
	data := ctx.Data

	// Handle case where collection is a slice of TEntity
	if collection, ok := data.collection.([]TEntity); ok {
		responses := make([]TModel, len(collection))
		for i, item := range collection {
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

	// Handle case where collection is a single TEntity
	if item, ok := data.collection.(*TEntity); ok {
		return ctx.FiberCtx.JSON(
			model.WebResponse[TModel]{
				Data: *ctx.Mapper.ModelToResponse(item),
			},
		)
	}

	return fiber.ErrInternalServerError
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
