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
	controller := newController[TEntity, TModel, *model.ListRequest](c.Log, c.Mapper)

	return wrapperPlural(
		ctx,
		controller,
		func(cp *CallbackParam[*model.ListRequest]) ([]TEntity, error) {
			requestParsed := new(model.ListRequest)
			if err := parseRequest(cp.fiberCtx, requestParsed); err != nil {
				return nil, err
			}
			request := model.ListRequest{Include: requestParsed.Include}

			return c.UseCase.List(cp.context, request)
		},
	)
}

func (c *CrudController[TEntity, TModel]) GetByID(ctx *fiber.Ctx) error {
	controller := newController[TEntity, TModel, *model.GetByIDRequest[int]](c.Log, c.Mapper)

	return wrapperPlural(
		ctx,
		controller,
		func(cp *CallbackParam[*model.GetByIDRequest[int]]) ([]TEntity, error) {
			requestParsed := new(model.GetByIDRequest[int])
			if err := parseRequest(cp.fiberCtx, requestParsed); err != nil {
				return nil, err
			}
			request := model.GetByIDRequest[int]{Include: requestParsed.Include, ID: requestParsed.ID}

			return c.UseCase.GetByID(cp.context, request)
		},
	)
}

func (c *CrudController[TEntity, TModel]) GetByIDs(ctx *fiber.Ctx) error {
	controller := newController[TEntity, TModel, *model.GetByIDRequest[[]int]](c.Log, c.Mapper)

	return wrapperPlural(
		ctx,
		controller,
		func(cp *CallbackParam[*model.GetByIDRequest[[]int]]) ([]TEntity, error) {
			requestParsed := new(model.GetByIDRequest[[]int])
			if err := parseRequest(cp.fiberCtx, requestParsed); err != nil {
				return nil, err
			}
			request := model.GetByIDRequest[[]int]{Include: requestParsed.Include, ID: requestParsed.ID}

			return c.UseCase.GetByIDs(cp.context, request)
		},
	)
}

func (c *CrudController[TEntity, TModel]) GetFirstByID(ctx *fiber.Ctx) error {
	controller := newController[TEntity, TModel, model.GetByIDRequest[int]](c.Log, c.Mapper)

	return wrapperSingular(
		ctx,
		controller,
		func(cp *CallbackParam[model.GetByIDRequest[int]]) (*TEntity, error) {
			requestParsed := new(model.GetByIDRequest[int])
			if err := parseRequest(cp.fiberCtx, requestParsed); err != nil {
				return nil, err
			}
			request := model.GetByIDRequest[int]{Include: requestParsed.Include, ID: requestParsed.ID}

			return c.UseCase.GetFirstByID(cp.context, request)
		},
	)
}
