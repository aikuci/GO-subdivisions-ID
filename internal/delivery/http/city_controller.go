package http

import (
	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/middleware/requestid"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/usecase"
	"github.com/gofiber/fiber/v2"

	"go.uber.org/zap"
)

type CityController struct {
	CrudController CrudController[model.CityResponse] // embedded

	Log     *zap.Logger
	UseCase usecase.CityUseCase
}

func NewCityController(log *zap.Logger, useCase *usecase.CityUseCase) *CityController {
	crudController := NewCrudController(log, useCase)

	return &CityController{
		CrudController: *crudController,

		Log:     log,
		UseCase: *useCase,
	}
}

func (c *CityController) GetByID(ctx *fiber.Ctx) error {
	userContext := requestid.SetContext(ctx.UserContext(), ctx)
	logger := c.Log.With(zap.String(string("requestid"), requestid.FromContext(userContext)))

	id, _ := ctx.ParamsInt("id")
	request := &model.GetByIDRequest{
		ID: id,
	}

	responses, err := c.UseCase.GetByIDReturnCollections(userContext, request)
	if err != nil {
		logger.Warn(err.Error())
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.CityResponse]{Data: responses})
}
