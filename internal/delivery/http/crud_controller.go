package http

import (
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CrudController[T any] struct {
	Log     *zap.Logger
	UseCase usecase.CruderUseCase[T]
}

func NewCrudController[T any](log *zap.Logger, useCase usecase.CruderUseCase[T]) *CrudController[T] {
	return &CrudController[T]{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *CrudController[T]) List(ctx *fiber.Ctx) error {
	controller := Prepare(ctx, c.Log)

	request := &model.ListRequest{}
	responses, err := c.UseCase.List(controller.context, request)
	if err != nil {
		controller.log.Warn(err.Error())
		return err
	}

	return ctx.JSON(model.WebResponse[[]T]{Data: responses})
}

func (c *CrudController[T]) GetByID(ctx *fiber.Ctx) error {
	controller := Prepare(ctx, c.Log)

	id, _ := ctx.ParamsInt("id")
	request := &model.GetByIDRequest[int]{ID: id}
	responses, err := c.UseCase.GetByID(controller.context, request)
	if err != nil {
		controller.log.Warn(err.Error())
		return err
	}

	return ctx.JSON(model.WebResponse[[]T]{Data: responses})
}

func (c *CrudController[T]) GetByIDs(ctx *fiber.Ctx) error {
	controller := Prepare(ctx, c.Log)

	// TODO: Collect ids
	request := &model.GetByIDRequest[[]int]{}
	responses, err := c.UseCase.GetByIDs(controller.context, request)
	if err != nil {
		controller.log.Warn(err.Error())
		return err
	}

	return ctx.JSON(model.WebResponse[[]T]{Data: responses})
}

func (c *CrudController[T]) GetFirstByID(ctx *fiber.Ctx) error {
	controller := Prepare(ctx, c.Log)

	id, _ := ctx.ParamsInt("id")
	request := &model.GetByIDRequest[int]{ID: id}
	response, err := c.UseCase.GetFirstByID(controller.context, request)
	if err != nil {
		controller.log.Warn(err.Error())
		return err
	}

	return ctx.JSON(model.WebResponse[*T]{Data: response})
}
