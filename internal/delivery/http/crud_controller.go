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
	controller := NewController[TEntity, TModel, *model.ListRequest](c.Log, c.Mapper)

	return WrapperPlural(ctx, controller, c.listFn)
}
func (c *CrudController[TEntity, TModel]) listFn(cp *CallbackParam[*model.ListRequest]) ([]TEntity, error) {
	request := &model.ListRequest{}

	return c.UseCase.List(cp.context, request)
}

func (c *CrudController[TEntity, TModel]) GetByID(ctx *fiber.Ctx) error {
	controller := NewController[TEntity, TModel, *model.GetByIDRequest[int]](c.Log, c.Mapper)

	return WrapperPlural(ctx, controller, c.getByIDFn)
}
func (c *CrudController[TEntity, TModel]) getByIDFn(cp *CallbackParam[*model.GetByIDRequest[int]]) ([]TEntity, error) {
	id, err := ParseId[int](cp.fiberCtx)
	if err != nil {
		return nil, err
	}
	request := &model.GetByIDRequest[int]{ID: *id}

	return c.UseCase.GetByID(cp.context, request)
}

func (c *CrudController[TEntity, TModel]) GetByIDs(ctx *fiber.Ctx) error {
	controller := NewController[TEntity, TModel, *model.GetByIDRequest[[]int]](c.Log, c.Mapper)

	return WrapperPlural(ctx, controller, c.getByIDsFn)
}
func (c *CrudController[TEntity, TModel]) getByIDsFn(cp *CallbackParam[*model.GetByIDRequest[[]int]]) ([]TEntity, error) {
	ids, err := ParseIds[[]int](cp.fiberCtx)
	if err != nil {
		return nil, err
	}
	request := &model.GetByIDRequest[[]int]{ID: *ids}

	return c.UseCase.GetByIDs(cp.context, request)
}

func (c *CrudController[TEntity, TModel]) GetFirstByID(ctx *fiber.Ctx) error {
	controller := NewController[TEntity, TModel, *model.GetByIDRequest[int]](c.Log, c.Mapper)

	return WrapperSingular(ctx, controller, c.getFirstByIDFn)
}
func (c *CrudController[TEntity, TModel]) getFirstByIDFn(cp *CallbackParam[*model.GetByIDRequest[int]]) (*TEntity, error) {
	id, err := ParseId[int](cp.fiberCtx)
	if err != nil {
		return nil, err
	}
	request := &model.GetByIDRequest[int]{ID: *id}

	return c.UseCase.GetFirstByID(cp.context, request)
}
