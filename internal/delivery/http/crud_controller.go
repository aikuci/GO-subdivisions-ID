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

func (c *CrudController[TEntity, TModel]) GetByID(ctx *fiber.Ctx) error {
	controller := newController[TEntity, TModel, model.GetByIDRequest[int]](c.Log, c.Mapper)

	return wrapperPlural(
		ctx,
		controller,
		func(cp *CallbackParam[model.GetByIDRequest[int]]) ([]TEntity, error) {
			return c.UseCase.GetByID(cp.context, cp.request)
		},
	)
}

func (c *CrudController[TEntity, TModel]) GetByIDs(ctx *fiber.Ctx) error {
	controller := newController[TEntity, TModel, model.GetByIDRequest[[]int]](c.Log, c.Mapper)

	return wrapperPlural(
		ctx,
		controller,
		func(cp *CallbackParam[model.GetByIDRequest[[]int]]) ([]TEntity, error) {
			return c.UseCase.GetByIDs(cp.context, cp.request)
		},
	)
}

func (c *CrudController[TEntity, TModel]) GetFirstByID(ctx *fiber.Ctx) error {
	controller := newController[TEntity, TModel, model.GetByIDRequest[int]](c.Log, c.Mapper)

	return wrapperSingular(
		ctx,
		controller,
		func(cp *CallbackParam[model.GetByIDRequest[int]]) (*TEntity, error) {
			return c.UseCase.GetFirstByID(cp.context, cp.request)
		},
	)
}
