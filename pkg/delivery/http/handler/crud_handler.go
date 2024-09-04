package handler

import (
	"github.com/aikuci/go-subdivisions-id/pkg/model"
	"github.com/aikuci/go-subdivisions-id/pkg/model/mapper"
	appusecase "github.com/aikuci/go-subdivisions-id/pkg/usecase"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CrudHandler[TEntity any, TModel any] struct {
	Log     *zap.Logger
	UseCase appusecase.CruderUseCase[TEntity]
	Mapper  mapper.CruderMapper[TEntity, TModel]
}

func NewCrudHandler[TEntity any, TModel any](log *zap.Logger, useCase appusecase.CruderUseCase[TEntity], mapper mapper.CruderMapper[TEntity, TModel]) *CrudHandler[TEntity, TModel] {
	return &CrudHandler[TEntity, TModel]{
		Log:     log,
		UseCase: useCase,
		Mapper:  mapper,
	}
}

func (c *CrudHandler[TEntity, TModel]) List(ctx *fiber.Ctx) error {
	return Wrapper(
		NewContext[model.ListRequest](c.Log, ctx, c.Mapper),
		func(ctx *HandlerContext[model.ListRequest, TEntity, TModel]) (any, int64, error) {
			return c.UseCase.List(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *CrudHandler[TEntity, TModel]) GetById(ctx *fiber.Ctx) error {
	return Wrapper(
		NewContext[model.GetByIDRequest[int]](c.Log, ctx, c.Mapper),
		func(ctx *HandlerContext[model.GetByIDRequest[int], TEntity, TModel]) (any, int64, error) {
			return c.UseCase.GetById(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *CrudHandler[TEntity, TModel]) GetFirstById(ctx *fiber.Ctx) error {
	return Wrapper(
		NewContext[model.GetByIDRequest[int]](c.Log, ctx, c.Mapper),
		func(ctx *HandlerContext[model.GetByIDRequest[int], TEntity, TModel]) (any, int64, error) {
			return c.UseCase.GetFirstById(ctx.Ctx, ctx.Request)
		},
	)
}
