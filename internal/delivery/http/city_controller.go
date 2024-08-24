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

func (c *CityController) GetByIdProvince(ctx *fiber.Ctx) error {
	controller := NewController[entity.City, model.CityResponse, *model.GetCityByIDProvinceRequest[int]](c.CrudController.Log, c.CrudController.Mapper)

	return WrapperPlural(
		ctx,
		controller,
		c.getByIDProvinceFn,
	)
}
func (c *CityController) getByIDProvinceFn(cp *CallbackParam[*model.GetCityByIDProvinceRequest[int]]) ([]entity.City, error) {
	id_province := ParseIntFromParamOrQuery(cp.fiberCtx, "id_province")
	request := &model.GetCityByIDProvinceRequest[int]{IDProvince: id_province}

	return c.UseCase.GetByIDProvince(cp.context, request)
}

func (c *CityController) GetByIDAndIdProvince(ctx *fiber.Ctx) error {
	controller := NewController[entity.City, model.CityResponse, *model.GetCityByIDRequest[int]](c.CrudController.Log, c.CrudController.Mapper)

	return WrapperSingular(
		ctx,
		controller,
		c.getByIDAndIdProvinceFn,
	)
}
func (c *CityController) getByIDAndIdProvinceFn(cp *CallbackParam[*model.GetCityByIDRequest[int]]) (*entity.City, error) {
	id, err := ParseId[int](cp.fiberCtx)
	if err != nil {
		return nil, err
	}
	id_province := ParseIntFromParamOrQuery(cp.fiberCtx, "id_province")
	request := &model.GetCityByIDRequest[int]{ID: *id, IDProvince: id_province}

	return c.UseCase.GetFirstByIDAndIDProvince(cp.context, request)
}
