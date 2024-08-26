package http

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/model/mapper"
	"github.com/aikuci/go-subdivisions-id/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type DistrictController struct {
	CrudController CrudController[entity.District, model.DistrictResponse] // embedded

	UseCase usecase.DistrictUseCase
}

func NewDistrictController(log *zap.Logger, useCase *usecase.DistrictUseCase, mapper mapper.CruderMapper[entity.District, model.DistrictResponse]) *DistrictController {
	crudController := NewCrudController(log, useCase, mapper)

	return &DistrictController{
		CrudController: *crudController,

		UseCase: *useCase,
	}
}

func (c *DistrictController) ListByIdAndIdCityAndIdProvince(ctx *fiber.Ctx) error {
	controller := newController[entity.District, model.DistrictResponse, model.ListDistrictByIDRequest[[]int]](c.CrudController.Log, c.CrudController.Mapper)

	return wrapperPlural(
		ctx,
		controller,
		func(cp *CallbackParam[model.ListDistrictByIDRequest[[]int]]) ([]entity.District, error) {
			return c.UseCase.ListFindByIdAndIdCityAndIdProvince(cp.context, cp.request)
		},
	)
}

func (c *DistrictController) GetByIdAndIdCityAndIdProvince(ctx *fiber.Ctx) error {
	controller := newController[entity.District, model.DistrictResponse, model.GetDistrictByIDRequest[[]int]](c.CrudController.Log, c.CrudController.Mapper)

	return wrapperPlural(
		ctx,
		controller,
		func(cp *CallbackParam[model.GetDistrictByIDRequest[[]int]]) ([]entity.District, error) {
			return c.UseCase.GetFindByIdAndIdCityAndIdProvince(cp.context, cp.request)
		},
	)
}

func (c *DistrictController) GetFirstByIdAndIdCityAndIdProvince(ctx *fiber.Ctx) error {
	controller := newController[entity.District, model.DistrictResponse, model.GetDistrictByIDRequest[int]](c.CrudController.Log, c.CrudController.Mapper)

	return wrapperSingular(
		ctx,
		controller,
		func(cp *CallbackParam[model.GetDistrictByIDRequest[int]]) (*entity.District, error) {
			return c.UseCase.GetFirstByIdAndIdCityAndIdProvince(cp.context, cp.request)
		},
	)
}
