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
	controller := NewController[entity.City, model.CityResponse, *model.ListCityByIDRequest[[]int]](c.CrudController.Log, c.CrudController.Mapper)

	return WrapperPlural(ctx, controller, c.listByIdAndIdProvinceFn)
}
func (c *CityController) listByIdAndIdProvinceFn(cp *CallbackParam[*model.ListCityByIDRequest[[]int]]) ([]entity.City, error) {
	requestParsed := new(model.ListCityByIDRequest[[]int])
	if err := ParseRequest(cp.fiberCtx, requestParsed); err != nil {
		return nil, err
	}
	request := &model.ListCityByIDRequest[[]int]{ID: requestParsed.ID, IDProvince: requestParsed.IDProvince}

	return c.UseCase.ListFindByIDAndIDProvince(cp.context, request)
}

func (c *CityController) GetByIDAndIdProvince(ctx *fiber.Ctx) error {
	controller := NewController[entity.City, model.CityResponse, *model.GetCityByIDRequest[[]int]](c.CrudController.Log, c.CrudController.Mapper)

	return WrapperPlural(ctx, controller, c.getByIDAndIdProvinceFn)
}
func (c *CityController) getByIDAndIdProvinceFn(cp *CallbackParam[*model.GetCityByIDRequest[[]int]]) ([]entity.City, error) {
	requestParsed := new(model.GetCityByIDRequest[[]int])
	if err := ParseRequest(cp.fiberCtx, requestParsed); err != nil {
		return nil, err
	}
	request := &model.GetCityByIDRequest[[]int]{GetByIDRequest: model.GetByIDRequest[[]int]{ID: requestParsed.ID}, IDProvince: requestParsed.IDProvince}

	return c.UseCase.GetFindByIDAndIDProvince(cp.context, request)
}

func (c *CityController) GetFirstByIDAndIdProvince(ctx *fiber.Ctx) error {
	controller := NewController[entity.City, model.CityResponse, *model.GetCityByIDRequest[int]](c.CrudController.Log, c.CrudController.Mapper)

	return WrapperSingular(ctx, controller, c.getFirstByIDAndIdProvinceFn)
}
func (c *CityController) getFirstByIDAndIdProvinceFn(cp *CallbackParam[*model.GetCityByIDRequest[int]]) (*entity.City, error) {
	id, err := ParseId[int](cp.fiberCtx)
	if err != nil {
		return nil, err
	}
	id_province := ParseIntFromParamOrQuery(cp.fiberCtx, "id_province")
	request := &model.GetCityByIDRequest[int]{GetByIDRequest: model.GetByIDRequest[int]{ID: *id}, IDProvince: id_province}

	return c.UseCase.GetFirstByIDAndIDProvince(cp.context, request)
}
