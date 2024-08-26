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
	Log      *zap.Logger
	FiberCtx *fiber.Ctx
	Mapper   mapper.CruderMapper[TEntity, TModel]
	Request  TRequest
}

func newController[TEntity any, TModel any, TRequest any](log *zap.Logger, fiberCtx *fiber.Ctx, mapper mapper.CruderMapper[TEntity, TModel]) *Controller[TEntity, TModel, TRequest] {
	return &Controller[TEntity, TModel, TRequest]{
		Log:      log,
		FiberCtx: fiberCtx,
		Mapper:   mapper,
	}
}

type CallbackArgs[T any] struct {
	context context.Context
	request T
}

func wrapperSingular[TEntity any, TModel any, TRequest any](c *Controller[TEntity, TModel, TRequest], callback func(ca *CallbackArgs[TRequest]) (*TEntity, error)) error {
	context := requestid.SetContext(c.FiberCtx.UserContext(), c.FiberCtx)
	log := c.Log.With(zap.String("requestid", requestid.FromContext(context)))

	requestParsed := new(TRequest)
	if err := parseRequest(c.FiberCtx, requestParsed); err != nil {
		return err
	}

	collection, err := callback(&CallbackArgs[TRequest]{context: context, request: *requestParsed})
	if err != nil {
		log.Warn(err.Error())
		return err
	}

	return c.FiberCtx.JSON(model.WebResponse[*TModel]{Data: c.Mapper.ModelToResponse(collection)})
}

func wrapperPlural[TEntity any, TModel any, TRequest any](c *Controller[TEntity, TModel, TRequest], callback func(ca *CallbackArgs[TRequest]) ([]TEntity, error)) error {
	context := requestid.SetContext(c.FiberCtx.UserContext(), c.FiberCtx)
	log := c.Log.With(zap.String("requestid", requestid.FromContext(context)))

	requestParsed := new(TRequest)
	if err := parseRequest(c.FiberCtx, requestParsed); err != nil {
		return err
	}

	collections, err := callback(&CallbackArgs[TRequest]{context: context, request: *requestParsed})
	if err != nil {
		log.Warn(err.Error())
		return err
	}

	responses := make([]TModel, len(collections))
	for i, collection := range collections {
		responses[i] = *c.Mapper.ModelToResponse(&collection)
	}

	return c.FiberCtx.JSON(model.WebResponse[[]TModel]{Data: responses})
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
