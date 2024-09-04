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

type District struct {
	UseCase usecase.District
	Mapper  appmapper.CruderMapper[entity.District, model.DistrictResponse]
}

func NewDistrict(useCase *usecase.District) *District {
	return &District{
		UseCase: *useCase,
		Mapper:  mapper.NewDistrict(),
	}
}

func (c *District) List(ctx *fiber.Ctx) error {
	return apphandler.Wrapper(
		apphandler.NewContext[model.ListDistrictByIDRequest[[]int]](ctx, c.Mapper),
		func(ctx *apphandler.Context[model.ListDistrictByIDRequest[[]int], entity.District, model.DistrictResponse]) (any, int64, error) {
			return c.UseCase.List(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *District) GetById(ctx *fiber.Ctx) error {
	return apphandler.Wrapper(
		apphandler.NewContext[model.GetDistrictByIDRequest[[]int]](ctx, c.Mapper),
		func(ctx *apphandler.Context[model.GetDistrictByIDRequest[[]int], entity.District, model.DistrictResponse]) (any, int64, error) {
			return c.UseCase.GetById(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *District) GetFirstById(ctx *fiber.Ctx) error {
	return apphandler.Wrapper(
		apphandler.NewContext[model.GetDistrictByIDRequest[int]](ctx, c.Mapper),
		func(ctx *apphandler.Context[model.GetDistrictByIDRequest[int], entity.District, model.DistrictResponse]) (any, int64, error) {
			return c.UseCase.GetFirstById(ctx.Ctx, ctx.Request)
		},
	)
}
