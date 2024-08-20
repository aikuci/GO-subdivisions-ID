package http

import (
	"fmt"

	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/middleware/requestid"
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
	userContext := requestid.SetContext(ctx.UserContext(), ctx)
	logger := c.Log.With(zap.String(string("requestid"), requestid.FromContext(userContext)))

	request := &model.ListProvinceRequest{}

	responses, err := c.UseCase.List(userContext, request)
	if err != nil {
		logger.Warn(err.Error())
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.ProvinceResponse]{Data: responses})
}

func (c *ProvinceController) Get(ctx *fiber.Ctx) error {
	userContext := requestid.SetContext(ctx.UserContext(), ctx)
	logger := c.Log.With(zap.String(string("requestid"), requestid.FromContext(userContext)))

	ID, err := ctx.ParamsInt("ID")
	if err != nil {
		message := fmt.Sprintf("failed to parse province ID: %v", ctx.Params("ID"))
		logger.Warn(err.Error(), zap.String("error", message))
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: message,
		}
	}
	request := &model.GetProvinceRequest{
		ID: ID,
	}

	response, err := c.UseCase.Get(userContext, request)
	if err != nil {
		logger.Warn(err.Error())
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ProvinceResponse]{Data: response})
}
