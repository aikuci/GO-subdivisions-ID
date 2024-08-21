package http

import (
	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/middleware/requestid"
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
	userContext := requestid.SetContext(ctx.UserContext(), ctx)
	logger := c.Log.With(zap.String(string("requestid"), requestid.FromContext(userContext)))

	request := &model.ListRequest{}

	responses, err := c.UseCase.List(userContext, request)
	if err != nil {
		logger.Warn(err.Error())
		return err
	}

	return ctx.JSON(model.WebResponse[[]T]{Data: responses})
}

func (c *CrudController[T]) GetByID(ctx *fiber.Ctx) error {
	userContext := requestid.SetContext(ctx.UserContext(), ctx)
	logger := c.Log.With(zap.String(string("requestid"), requestid.FromContext(userContext)))

	id, _ := ctx.ParamsInt("id")
	request := &model.GetByIDRequest{
		ID: id,
	}

	response, err := c.UseCase.GetByID(userContext, request)
	if err != nil {
		logger.Warn(err.Error())
		return err
	}

	return ctx.JSON(model.WebResponse[*T]{Data: response})
}
