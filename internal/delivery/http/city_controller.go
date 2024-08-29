package http

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/usecase"
	apphttp "github.com/aikuci/go-subdivisions-id/pkg/delivery/http"
	appmapper "github.com/aikuci/go-subdivisions-id/pkg/model/mapper"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CityController struct {
	Log     *zap.Logger
	UseCase usecase.CityUseCase
	Mapper  appmapper.CruderMapper[entity.City, model.CityResponse]
}

func NewCityController(log *zap.Logger, useCase *usecase.CityUseCase, mapper appmapper.CruderMapper[entity.City, model.CityResponse]) *CityController {
	return &CityController{
		Log:     log,
		UseCase: *useCase,
		Mapper:  mapper,
	}
}

func (c *CityController) List(ctx *fiber.Ctx) error {
	return apphttp.Wrapper(
		apphttp.NewContext[model.ListCityByIDRequest[[]int]](c.Log, ctx, c.Mapper),
		func(ctx *apphttp.ControllerContext[model.ListCityByIDRequest[[]int], entity.City, model.CityResponse]) (any, int64, error) {
			return c.UseCase.List(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *CityController) GetById(ctx *fiber.Ctx) error {
	return apphttp.Wrapper(
		apphttp.NewContext[model.GetCityByIDRequest[[]int]](c.Log, ctx, c.Mapper),
		func(ctx *apphttp.ControllerContext[model.GetCityByIDRequest[[]int], entity.City, model.CityResponse]) (any, int64, error) {
			return c.UseCase.GetById(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *CityController) GetFirstById(ctx *fiber.Ctx) error {
	return apphttp.Wrapper(
		apphttp.NewContext[model.GetCityByIDRequest[int]](c.Log, ctx, c.Mapper),
		func(ctx *apphttp.ControllerContext[model.GetCityByIDRequest[int], entity.City, model.CityResponse]) (any, int64, error) {
			return c.UseCase.GetFirstById(ctx.Ctx, ctx.Request)
		},
	)
}
