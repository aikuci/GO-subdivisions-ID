package http

import (
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

	request := &model.ListRequest{}

	responses, err := c.UseCase.List(userContext, request)
	if err != nil {
		logger.Warn(err.Error())
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.ProvinceResponse]{Data: responses})
}

func (c *ProvinceController) GetByID(ctx *fiber.Ctx) error {
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

	return ctx.JSON(model.WebResponse[*model.ProvinceResponse]{Data: response})
}
