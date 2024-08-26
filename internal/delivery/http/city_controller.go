package http

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/model/mapper"
	"github.com/aikuci/go-subdivisions-id/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CityController struct {
	Log     *zap.Logger
	UseCase usecase.CityUseCase
	Mapper  mapper.CruderMapper[entity.City, model.CityResponse]
}

func NewCityController(log *zap.Logger, useCase *usecase.CityUseCase, mapper mapper.CruderMapper[entity.City, model.CityResponse]) *CityController {
	return &CityController{
		Log:     log,
		UseCase: *useCase,
		Mapper:  mapper,
	}
}

func (c *CityController) List(ctx *fiber.Ctx) error {
	return wrapperPlural(
		newController[entity.City, model.CityResponse, model.ListCityByIDRequest[[]int]](c.Log, ctx, c.Mapper),
		func(ca *CallbackArgs[model.ListCityByIDRequest[[]int]]) ([]entity.City, error) {
			return c.UseCase.List(ca.context, ca.request)
		},
	)
}

func (c *CityController) GetById(ctx *fiber.Ctx) error {
	return wrapperPlural(
		newController[entity.City, model.CityResponse, model.GetCityByIDRequest[[]int]](c.Log, ctx, c.Mapper),
		func(ca *CallbackArgs[model.GetCityByIDRequest[[]int]]) ([]entity.City, error) {
			return c.UseCase.GetById(ca.context, ca.request)
		},
	)
}

func (c *CityController) GetFirstById(ctx *fiber.Ctx) error {
	return wrapperSingular(
		newController[entity.City, model.CityResponse, model.GetCityByIDRequest[int]](c.Log, ctx, c.Mapper),
		func(ca *CallbackArgs[model.GetCityByIDRequest[int]]) (*entity.City, error) {
			return c.UseCase.GetFirstById(ca.context, ca.request)
		},
	)
}
