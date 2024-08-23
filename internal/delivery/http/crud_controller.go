package http

import (
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CrudController[T any] struct {
	Controller *Controller

	UseCase usecase.CruderUseCase[T]
}

func NewCrudController[T any](log *zap.Logger, useCase usecase.CruderUseCase[T]) *CrudController[T] {
	controller := NewController(log)

	return &CrudController[T]{
		Controller: controller,

		UseCase: useCase,
	}
}

func (c *CrudController[T]) List(ctx *fiber.Ctx) error {
	c.Controller.Prepare(ctx)

	request := &model.ListRequest{}
	responses, err := c.UseCase.List(c.Controller.context, request)
	if err != nil {
		c.Controller.Log.Warn(err.Error())
		return err
	}

	return ctx.JSON(model.WebResponse[[]T]{Data: responses})
}

func (c *CrudController[T]) GetByID(ctx *fiber.Ctx) error {
	c.Controller.Prepare(ctx)

	id, _ := ctx.ParamsInt("id")
	request := &model.GetByIDRequest[int]{ID: id}
	responses, err := c.UseCase.GetByID(c.Controller.context, request)
	if err != nil {
		c.Controller.Log.Warn(err.Error())
		return err
	}

	return ctx.JSON(model.WebResponse[[]T]{Data: responses})
}

func (c *CrudController[T]) GetByIDs(ctx *fiber.Ctx) error {
	c.Controller.Prepare(ctx)

	// TODO: Collect ids
	request := &model.GetByIDRequest[[]int]{}
	responses, err := c.UseCase.GetByIDs(c.Controller.context, request)
	if err != nil {
		c.Controller.Log.Warn(err.Error())
		return err
	}

	return ctx.JSON(model.WebResponse[[]T]{Data: responses})
}

func (c *CrudController[T]) GetFirstByID(ctx *fiber.Ctx) error {
	c.Controller.Prepare(ctx)

	id, _ := ctx.ParamsInt("id")
	request := &model.GetByIDRequest[int]{ID: id}
	response, err := c.UseCase.GetFirstByID(c.Controller.context, request)
	if err != nil {
		c.Controller.Log.Warn(err.Error())
		return err
	}

	return ctx.JSON(model.WebResponse[*T]{Data: response})
}
