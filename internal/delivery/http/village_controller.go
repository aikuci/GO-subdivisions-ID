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

type VillageController struct {
	Log     *zap.Logger
	UseCase usecase.VillageUseCase
	Mapper  appmapper.CruderMapper[entity.Village, model.VillageResponse]
}

func NewVillageController(log *zap.Logger, useCase *usecase.VillageUseCase, mapper appmapper.CruderMapper[entity.Village, model.VillageResponse]) *VillageController {
	return &VillageController{
		Log:     log,
		UseCase: *useCase,
		Mapper:  mapper,
	}
}

func (c *VillageController) List(ctx *fiber.Ctx) error {
	return apphttp.Wrapper(
		apphttp.NewContext[model.ListVillageByIDRequest[[]int]](c.Log, ctx, c.Mapper),
		func(ctx *apphttp.ControllerContext[model.ListVillageByIDRequest[[]int], entity.Village, model.VillageResponse]) (any, int64, error) {
			return c.UseCase.List(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *VillageController) GetById(ctx *fiber.Ctx) error {
	return apphttp.Wrapper(
		apphttp.NewContext[model.GetVillageByIDRequest[[]int]](c.Log, ctx, c.Mapper),
		func(ctx *apphttp.ControllerContext[model.GetVillageByIDRequest[[]int], entity.Village, model.VillageResponse]) (any, int64, error) {
			return c.UseCase.GetById(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *VillageController) GetFirstById(ctx *fiber.Ctx) error {
	return apphttp.Wrapper(
		apphttp.NewContext[model.GetVillageByIDRequest[int]](c.Log, ctx, c.Mapper),
		func(ctx *apphttp.ControllerContext[model.GetVillageByIDRequest[int], entity.Village, model.VillageResponse]) (any, int64, error) {
			return c.UseCase.GetFirstById(ctx.Ctx, ctx.Request)
		},
	)
}
