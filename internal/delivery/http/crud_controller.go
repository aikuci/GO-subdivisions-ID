package http

import (
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/model/mapper"
	"github.com/aikuci/go-subdivisions-id/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CrudController[TEntity any, TModel any, TRelation []string] struct {
	Log     *zap.Logger
	UseCase usecase.CruderUseCase[TEntity, TRelation]
	Mapper  mapper.CruderMapper[TEntity, TModel]
}

func NewCrudController[TEntity any, TModel any, TRelation []string](log *zap.Logger, useCase usecase.CruderUseCase[TEntity, TRelation], mapper mapper.CruderMapper[TEntity, TModel]) *CrudController[TEntity, TModel, TRelation] {
	return &CrudController[TEntity, TModel, TRelation]{
		Log:     log,
		UseCase: useCase,
		Mapper:  mapper,
	}
}

func (c *CrudController[TEntity, TModel, TRelation]) List(ctx *fiber.Ctx) error {
	controller := NewController[TEntity, TModel, *model.ListRequest](c.Log, c.Mapper)

	return WrapperPlural(ctx, controller, c.listFn)
}
func (c *CrudController[TEntity, TModel, TRelation]) listFn(cp *CallbackParam[*model.ListRequest]) ([]TEntity, error) {
	requestParsed := new(model.ListRequest)
	if err := ParseRequest(cp.fiberCtx, requestParsed); err != nil {
		return nil, err
	}
	request := &model.ListRequest{Include: requestParsed.Include}

	return c.UseCase.List(cp.context, request)
}

func (c *CrudController[TEntity, TModel, TRelation]) GetByID(ctx *fiber.Ctx) error {
	controller := NewController[TEntity, TModel, *model.GetByIDRequest[int]](c.Log, c.Mapper)

	return WrapperPlural(ctx, controller, c.getByIDFn)
}
func (c *CrudController[TEntity, TModel, TRelation]) getByIDFn(cp *CallbackParam[*model.GetByIDRequest[int]]) ([]TEntity, error) {
	requestParsed := new(model.GetByIDRequest[int])
	if err := ParseRequest(cp.fiberCtx, requestParsed); err != nil {
		return nil, err
	}
	request := &model.GetByIDRequest[int]{Include: requestParsed.Include, ID: requestParsed.ID}

	return c.UseCase.GetByID(cp.context, request)
}

func (c *CrudController[TEntity, TModel, TRelation]) GetByIDs(ctx *fiber.Ctx) error {
	controller := NewController[TEntity, TModel, *model.GetByIDRequest[[]int]](c.Log, c.Mapper)

	return WrapperPlural(ctx, controller, c.getByIDsFn)
}
func (c *CrudController[TEntity, TModel, TRelation]) getByIDsFn(cp *CallbackParam[*model.GetByIDRequest[[]int]]) ([]TEntity, error) {
	requestParsed := new(model.GetByIDRequest[[]int])
	if err := ParseRequest(cp.fiberCtx, requestParsed); err != nil {
		return nil, err
	}
	request := &model.GetByIDRequest[[]int]{Include: requestParsed.Include, ID: requestParsed.ID}

	return c.UseCase.GetByIDs(cp.context, request)
}

func (c *CrudController[TEntity, TModel, TRelation]) GetFirstByID(ctx *fiber.Ctx) error {
	controller := NewController[TEntity, TModel, *model.GetByIDRequest[int]](c.Log, c.Mapper)

	return WrapperSingular(ctx, controller, c.getFirstByIDFn)
}
func (c *CrudController[TEntity, TModel, TRelation]) getFirstByIDFn(cp *CallbackParam[*model.GetByIDRequest[int]]) (*TEntity, error) {
	requestParsed := new(model.GetByIDRequest[int])
	if err := ParseRequest(cp.fiberCtx, requestParsed); err != nil {
		return nil, err
	}
	request := &model.GetByIDRequest[int]{Include: requestParsed.Include, ID: requestParsed.ID}

	return c.UseCase.GetFirstByID(cp.context, request)
}
