package http

import (
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/model/mapper"
	"github.com/aikuci/go-subdivisions-id/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CrudController[TEntity any, TModel any] struct {
	Log     *zap.Logger
	UseCase usecase.CruderUseCase[TEntity]
	Mapper  mapper.CruderMapper[TEntity, TModel]
}

func NewCrudController[TEntity any, TModel any](log *zap.Logger, useCase usecase.CruderUseCase[TEntity], mapper mapper.CruderMapper[TEntity, TModel]) *CrudController[TEntity, TModel] {
	return &CrudController[TEntity, TModel]{
		Log:     log,
		UseCase: useCase,
		Mapper:  mapper,
	}
}

func (c *CrudController[TEntity, TModel]) List(ctx *fiber.Ctx) error {
	controller := newController[TEntity, TModel, model.ListRequest](c.Log, c.Mapper)

	return wrapperPlural(
		ctx,
		controller,
		func(cp *CallbackParam[model.ListRequest]) ([]TEntity, error) {
			return c.UseCase.List(cp.context, cp.request)
		},
	)
}

func (c *CrudController[TEntity, TModel]) GetById(ctx *fiber.Ctx) error {
	controller := newController[TEntity, TModel, model.GetByIDRequest[int]](c.Log, c.Mapper)

	return wrapperPlural(
		ctx,
		controller,
		func(cp *CallbackParam[model.GetByIDRequest[int]]) ([]TEntity, error) {
			return c.UseCase.GetById(cp.context, cp.request)
		},
	)
}

func (c *CrudController[TEntity, TModel]) GetByIds(ctx *fiber.Ctx) error {
	controller := newController[TEntity, TModel, model.GetByIDRequest[[]int]](c.Log, c.Mapper)

	return wrapperPlural(
		ctx,
		controller,
		func(cp *CallbackParam[model.GetByIDRequest[[]int]]) ([]TEntity, error) {
			return c.UseCase.GetByIds(cp.context, cp.request)
		},
	)
}

func (c *CrudController[TEntity, TModel]) GetFirstById(ctx *fiber.Ctx) error {
	controller := newController[TEntity, TModel, model.GetByIDRequest[int]](c.Log, c.Mapper)

	return wrapperSingular(
		ctx,
		controller,
		func(cp *CallbackParam[model.GetByIDRequest[int]]) (*TEntity, error) {
			return c.UseCase.GetFirstById(cp.context, cp.request)
		},
	)
}
