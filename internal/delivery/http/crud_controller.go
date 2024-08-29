package http

import (
	appmodel "github.com/aikuci/go-subdivisions-id/pkg/model"
	appmapper "github.com/aikuci/go-subdivisions-id/pkg/model/mapper"
	appusecase "github.com/aikuci/go-subdivisions-id/pkg/usecase"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CrudController[TEntity any, TModel any] struct {
	Log     *zap.Logger
	UseCase appusecase.CruderUseCase[TEntity]
	Mapper  appmapper.CruderMapper[TEntity, TModel]
}

func NewCrudController[TEntity any, TModel any](log *zap.Logger, useCase appusecase.CruderUseCase[TEntity], mapper appmapper.CruderMapper[TEntity, TModel]) *CrudController[TEntity, TModel] {
	return &CrudController[TEntity, TModel]{
		Log:     log,
		UseCase: useCase,
		Mapper:  mapper,
	}
}

func (c *CrudController[TEntity, TModel]) List(ctx *fiber.Ctx) error {
	return Wrapper(
		NewContext[appmodel.ListRequest](c.Log, ctx, c.Mapper),
		func(ctx *ControllerContext[appmodel.ListRequest, TEntity, TModel]) (any, int64, error) {
			return c.UseCase.List(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *CrudController[TEntity, TModel]) GetById(ctx *fiber.Ctx) error {
	return Wrapper(
		NewContext[appmodel.GetByIDRequest[int]](c.Log, ctx, c.Mapper),
		func(ctx *ControllerContext[appmodel.GetByIDRequest[int], TEntity, TModel]) (any, int64, error) {
			return c.UseCase.GetById(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *CrudController[TEntity, TModel]) GetFirstById(ctx *fiber.Ctx) error {
	return Wrapper(
		NewContext[appmodel.GetByIDRequest[int]](c.Log, ctx, c.Mapper),
		func(ctx *ControllerContext[appmodel.GetByIDRequest[int], TEntity, TModel]) (any, int64, error) {
			return c.UseCase.GetFirstById(ctx.Ctx, ctx.Request)
		},
	)
}
