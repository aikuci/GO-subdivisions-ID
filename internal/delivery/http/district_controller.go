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

type DistrictController struct {
	Log     *zap.Logger
	UseCase usecase.DistrictUseCase
	Mapper  appmapper.CruderMapper[entity.District, model.DistrictResponse]
}

func NewDistrictController(log *zap.Logger, useCase *usecase.DistrictUseCase, mapper appmapper.CruderMapper[entity.District, model.DistrictResponse]) *DistrictController {
	return &DistrictController{
		Log:     log,
		UseCase: *useCase,
		Mapper:  mapper,
	}
}

func (c *DistrictController) List(ctx *fiber.Ctx) error {
	return apphttp.Wrapper(
		apphttp.NewContext[model.ListDistrictByIDRequest[[]int]](c.Log, ctx, c.Mapper),
		func(ctx *apphttp.ControllerContext[model.ListDistrictByIDRequest[[]int], entity.District, model.DistrictResponse]) (any, int64, error) {
			return c.UseCase.List(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *DistrictController) GetById(ctx *fiber.Ctx) error {
	return apphttp.Wrapper(
		apphttp.NewContext[model.GetDistrictByIDRequest[[]int]](c.Log, ctx, c.Mapper),
		func(ctx *apphttp.ControllerContext[model.GetDistrictByIDRequest[[]int], entity.District, model.DistrictResponse]) (any, int64, error) {
			return c.UseCase.GetById(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *DistrictController) GetFirstById(ctx *fiber.Ctx) error {
	return apphttp.Wrapper(
		apphttp.NewContext[model.GetDistrictByIDRequest[int]](c.Log, ctx, c.Mapper),
		func(ctx *apphttp.ControllerContext[model.GetDistrictByIDRequest[int], entity.District, model.DistrictResponse]) (any, int64, error) {
			return c.UseCase.GetFirstById(ctx.Ctx, ctx.Request)
		},
	)
}
