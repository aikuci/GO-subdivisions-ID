package http

import (
	"context"

	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/middleware/requestid"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/model/mapper"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Controller[TEntity any, TModel any, TRequest any] struct {
	Log     *zap.Logger
	Mapper  mapper.CruderMapper[TEntity, TModel]
	Request TRequest
}

func NewController[TEntity any, TModel any, TRequest any](log *zap.Logger, mapper mapper.CruderMapper[TEntity, TModel]) *Controller[TEntity, TModel, TRequest] {
	return &Controller[TEntity, TModel, TRequest]{
		Log:    log,
		Mapper: mapper,
	}
}

type CallbackParam[T any] struct {
	context  context.Context
	log      *zap.Logger
	fiberCtx *fiber.Ctx
	request  T
}

func WrapperSingular[TEntity any, TModel any, TRequest any](ctx *fiber.Ctx, c *Controller[TEntity, TModel, TRequest], callback func(cp *CallbackParam[TRequest]) (*TEntity, error)) error {
	context := requestid.SetContext(ctx.UserContext(), ctx)
	log := c.Log.With(zap.String("requestid", requestid.FromContext(context)))

	collection, err := callback(&CallbackParam[TRequest]{context: context, log: log, fiberCtx: ctx, request: c.Request})
	if err != nil {
		log.Warn(err.Error())
		return err
	}

	return ctx.JSON(model.WebResponse[*TModel]{Data: c.Mapper.ModelToResponse(collection)})
}

func WrapperPlural[TEntity any, TModel any, TRequest any](ctx *fiber.Ctx, c *Controller[TEntity, TModel, TRequest], callback func(cp *CallbackParam[TRequest]) ([]TEntity, error)) error {
	context := requestid.SetContext(ctx.UserContext(), ctx)
	log := c.Log.With(zap.String("requestid", requestid.FromContext(context)))

	collections, err := callback(&CallbackParam[TRequest]{context: context, log: log, fiberCtx: ctx, request: c.Request})
	if err != nil {
		log.Warn(err.Error())
		return err
	}

	responses := make([]TModel, len(collections))
	for i, collection := range collections {
		responses[i] = *c.Mapper.ModelToResponse(&collection)
	}

	return ctx.JSON(model.WebResponse[[]TModel]{Data: responses})
}

func ParseId[T model.IdSingular](ctx *fiber.Ctx) (*T, error) {
	id := new(model.IdSingularRequest[T])
	if err := ParseRequest(ctx, id); err != nil {
		return nil, err
	}
	return &id.ID, nil
}

func ParseIds[T model.IdPlural](ctx *fiber.Ctx) (*T, error) {
	ids := new(model.IdPluralRequest[T])
	if err := ParseRequest(ctx, ids); err != nil {
		return nil, err
	}
	return &ids.ID, nil
}

func ParseRequest(ctx *fiber.Ctx, request any) error {
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

func ParseIntFromParamOrQuery(ctx *fiber.Ctx, requestKey string) int {
	requestValue, err := ctx.ParamsInt(requestKey)
	if err == nil {
		return requestValue
	}

	return ctx.QueryInt(requestKey)
}
