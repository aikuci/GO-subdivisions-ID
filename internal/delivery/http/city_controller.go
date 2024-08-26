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
	CrudController CrudController[entity.City, model.CityResponse] // embedded

	UseCase usecase.CityUseCase
}

func NewCityController(log *zap.Logger, useCase *usecase.CityUseCase, mapper mapper.CruderMapper[entity.City, model.CityResponse]) *CityController {
	crudController := NewCrudController(log, useCase, mapper)

	return &CityController{
		CrudController: *crudController,

		UseCase: *useCase,
	}
}

func (c *CityController) ListByIdAndIdProvince(ctx *fiber.Ctx) error {
	controller := newController[entity.City, model.CityResponse, model.ListCityByIDRequest[[]int]](c.CrudController.Log, c.CrudController.Mapper)

	return wrapperPlural(
		ctx,
		controller,
		func(cp *CallbackParam[model.ListCityByIDRequest[[]int]]) ([]entity.City, error) {
			return c.UseCase.ListFindByIdAndIdProvince(cp.context, cp.request)
		},
	)
}

func (c *CityController) GetByIdAndIdProvince(ctx *fiber.Ctx) error {
	controller := newController[entity.City, model.CityResponse, model.GetCityByIDRequest[[]int]](c.CrudController.Log, c.CrudController.Mapper)

	return wrapperPlural(
		ctx,
		controller,
		func(cp *CallbackParam[model.GetCityByIDRequest[[]int]]) ([]entity.City, error) {
			return c.UseCase.GetFindByIdAndIdProvince(cp.context, cp.request)
		},
	)
}

func (c *CityController) GetFirstByIdAndIdProvince(ctx *fiber.Ctx) error {
	controller := newController[entity.City, model.CityResponse, model.GetCityByIDRequest[int]](c.CrudController.Log, c.CrudController.Mapper)

	return wrapperSingular(
		ctx,
		controller,
		func(cp *CallbackParam[model.GetCityByIDRequest[int]]) (*entity.City, error) {
			return c.UseCase.GetFirstByIdAndIdProvince(cp.context, cp.request)
		},
	)
}
