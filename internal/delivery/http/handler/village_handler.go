package handler

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/usecase"
	apphandler "github.com/aikuci/go-subdivisions-id/pkg/delivery/http/handler"
	appmapper "github.com/aikuci/go-subdivisions-id/pkg/model/mapper"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type VillageHandler struct {
	Log     *zap.Logger
	UseCase usecase.VillageUseCase
	Mapper  appmapper.CruderMapper[entity.Village, model.VillageResponse]
}

func NewVillageHandler(log *zap.Logger, useCase *usecase.VillageUseCase, mapper appmapper.CruderMapper[entity.Village, model.VillageResponse]) *VillageHandler {
	return &VillageHandler{
		Log:     log,
		UseCase: *useCase,
		Mapper:  mapper,
	}
}

func (c *VillageHandler) List(ctx *fiber.Ctx) error {
	return apphandler.Wrapper(
		apphandler.NewContext[model.ListVillageByIDRequest[[]int]](c.Log, ctx, c.Mapper),
		func(ctx *apphandler.HandlerContext[model.ListVillageByIDRequest[[]int], entity.Village, model.VillageResponse]) (any, int64, error) {
			return c.UseCase.List(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *VillageHandler) GetById(ctx *fiber.Ctx) error {
	return apphandler.Wrapper(
		apphandler.NewContext[model.GetVillageByIDRequest[[]int]](c.Log, ctx, c.Mapper),
		func(ctx *apphandler.HandlerContext[model.GetVillageByIDRequest[[]int], entity.Village, model.VillageResponse]) (any, int64, error) {
			return c.UseCase.GetById(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *VillageHandler) GetFirstById(ctx *fiber.Ctx) error {
	return apphandler.Wrapper(
		apphandler.NewContext[model.GetVillageByIDRequest[int]](c.Log, ctx, c.Mapper),
		func(ctx *apphandler.HandlerContext[model.GetVillageByIDRequest[int], entity.Village, model.VillageResponse]) (any, int64, error) {
			return c.UseCase.GetFirstById(ctx.Ctx, ctx.Request)
		},
	)
}
