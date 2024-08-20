package http

import (
	"fmt"

	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ProvinceController struct {
	Log     *zap.Logger
	UseCase *usecase.ProvinceUseCase
}

func NewProvinceController(log *zap.Logger, useCase *usecase.ProvinceUseCase) *ProvinceController {
	return &ProvinceController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *ProvinceController) List(ctx *fiber.Ctx) error {
	logger := c.Log
	if rid, ok := ctx.Locals("requestid").(string); ok {
		logger = logger.With(zap.String("requestid", rid))
	}

	request := &model.ListProvinceRequest{}

	ctx.SetUserContext(ctx.Context())
	responses, err := c.UseCase.List(ctx.UserContext(), request)
	if e, ok := err.(*fiber.Error); ok {
		logger.Warn(err.Error())
		return &fiber.Error{
			Code:    e.Code,
			Message: err.Error(),
		}
	}

	return ctx.JSON(model.WebResponse[[]model.ProvinceResponse]{Data: responses})
}

func (c *ProvinceController) Get(ctx *fiber.Ctx) error {
	logger := c.Log
	if rid, ok := ctx.Locals("requestid").(string); ok {
		logger = logger.With(zap.String("requestid", rid))
	}

	ID, err := ctx.ParamsInt("ID")
	if err != nil {
		message := fmt.Sprintf("failed to parse province ID %v", ID)
		logger.Warn(err.Error(), zap.String("error", message))
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: message,
		}
	}

	request := &model.GetProvinceRequest{
		ID: ID,
	}

	ctx.SetUserContext(ctx.Context())
	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if e, ok := err.(*fiber.Error); ok {
		logger.Warn(err.Error())
		return &fiber.Error{
			Code:    e.Code,
			Message: err.Error(),
		}
	}

	return ctx.JSON(model.WebResponse[*model.ProvinceResponse]{Data: response})
}
