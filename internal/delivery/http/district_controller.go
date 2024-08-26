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
	Log     *zap.Logger
	UseCase usecase.DistrictUseCase
	Mapper  mapper.CruderMapper[entity.District, model.DistrictResponse]
}

func NewDistrictController(log *zap.Logger, useCase *usecase.DistrictUseCase, mapper mapper.CruderMapper[entity.District, model.DistrictResponse]) *DistrictController {
	return &DistrictController{
		Log:     log,
		UseCase: *useCase,
		Mapper:  mapper,
	}
}

func (c *DistrictController) List(ctx *fiber.Ctx) error {
	return wrapperPlural(
		newController[entity.District, model.DistrictResponse, model.ListDistrictByIDRequest[[]int]](c.Log, ctx, c.Mapper),
		func(ca *CallbackArgs[model.ListDistrictByIDRequest[[]int]]) ([]entity.District, error) {
			return c.UseCase.List(ca.context, ca.request)
		},
	)
}

func (c *DistrictController) GetById(ctx *fiber.Ctx) error {

	return wrapperPlural(
		newController[entity.District, model.DistrictResponse, model.GetDistrictByIDRequest[[]int]](c.Log, ctx, c.Mapper),
		func(ca *CallbackArgs[model.GetDistrictByIDRequest[[]int]]) ([]entity.District, error) {
			return c.UseCase.GetById(ca.context, ca.request)
		},
	)
}

func (c *DistrictController) GetFirstById(ctx *fiber.Ctx) error {
	return wrapperSingular(
		newController[entity.District, model.DistrictResponse, model.GetDistrictByIDRequest[int]](c.Log, ctx, c.Mapper),
		func(ca *CallbackArgs[model.GetDistrictByIDRequest[int]]) (*entity.District, error) {
			return c.UseCase.GetFirstById(ca.context, ca.request)
		},
	)
}
