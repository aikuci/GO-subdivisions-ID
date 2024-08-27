package http

import (
	"github.com/aikuci/go-subdivisions-id/internal/usecase"
	appmodel "github.com/aikuci/go-subdivisions-id/pkg/model"
	appmapper "github.com/aikuci/go-subdivisions-id/pkg/model/mapper"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CrudController[TEntity any, TModel any] struct {
	Log     *zap.Logger
	UseCase usecase.CruderUseCase[TEntity]
	Mapper  appmapper.CruderMapper[TEntity, TModel]
}

func NewCrudController[TEntity any, TModel any](log *zap.Logger, useCase usecase.CruderUseCase[TEntity], mapper appmapper.CruderMapper[TEntity, TModel]) *CrudController[TEntity, TModel] {
	return &CrudController[TEntity, TModel]{
		Log:     log,
		UseCase: useCase,
		Mapper:  mapper,
	}
}

func (c *CrudController[TEntity, TModel]) List(ctx *fiber.Ctx) error {
	return wrapperPlural(
		newController[TEntity, TModel, appmodel.ListRequest](c.Log, ctx, c.Mapper),
		func(ca *CallbackArgs[appmodel.ListRequest]) ([]TEntity, int64, error) {
			return c.UseCase.List(ca.context, ca.request)
		},
	)
}

func (c *CrudController[TEntity, TModel]) GetById(ctx *fiber.Ctx) error {
	return wrapperPlural(
		newController[TEntity, TModel, appmodel.GetByIDRequest[int]](c.Log, ctx, c.Mapper),
		func(ca *CallbackArgs[appmodel.GetByIDRequest[int]]) ([]TEntity, int64, error) {
			return c.UseCase.GetById(ca.context, ca.request)
		},
	)
}

func (c *CrudController[TEntity, TModel]) GetFirstById(ctx *fiber.Ctx) error {
	return wrapperSingular(
		newController[TEntity, TModel, appmodel.GetByIDRequest[int]](c.Log, ctx, c.Mapper),
		func(ca *CallbackArgs[appmodel.GetByIDRequest[int]]) (*TEntity, error) {
			return c.UseCase.GetFirstById(ca.context, ca.request)
		},
	)
}
