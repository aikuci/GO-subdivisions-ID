package handler

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/model/mapper"
	"github.com/aikuci/go-subdivisions-id/internal/usecase"
	apphandler "github.com/aikuci/go-subdivisions-id/pkg/delivery/http/handler"
	appmapper "github.com/aikuci/go-subdivisions-id/pkg/model/mapper"

	"github.com/gofiber/fiber/v2"
)

type Village struct {
	UseCase usecase.Village
	Mapper  appmapper.CruderMapper[entity.Village, model.VillageResponse]
}

func NewVillage(useCase *usecase.Village) *Village {
	return &Village{
		UseCase: *useCase,
		Mapper:  mapper.NewVillage(),
	}
}

func (c *Village) List(ctx *fiber.Ctx) error {
	return apphandler.Wrapper(
		apphandler.NewContext[model.ListVillageByIDRequest[[]int]](ctx, c.Mapper),
		func(ctx *apphandler.Context[model.ListVillageByIDRequest[[]int], entity.Village, model.VillageResponse]) (any, int64, error) {
			return c.UseCase.List(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *Village) GetById(ctx *fiber.Ctx) error {
	return apphandler.Wrapper(
		apphandler.NewContext[model.GetVillageByIDRequest[[]int]](ctx, c.Mapper),
		func(ctx *apphandler.Context[model.GetVillageByIDRequest[[]int], entity.Village, model.VillageResponse]) (any, int64, error) {
			return c.UseCase.GetById(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *Village) GetFirstById(ctx *fiber.Ctx) error {
	return apphandler.Wrapper(
		apphandler.NewContext[model.GetVillageByIDRequest[int]](ctx, c.Mapper),
		func(ctx *apphandler.Context[model.GetVillageByIDRequest[int], entity.Village, model.VillageResponse]) (any, int64, error) {
			return c.UseCase.GetFirstById(ctx.Ctx, ctx.Request)
		},
	)
}
