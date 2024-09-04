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

type City struct {
	UseCase usecase.City
	Mapper  appmapper.CruderMapper[entity.City, model.CityResponse]
}

func NewCity(useCase *usecase.City) *City {
	return &City{
		UseCase: *useCase,
		Mapper:  mapper.NewCity(),
	}
}

func (c *City) List(ctx *fiber.Ctx) error {
	return apphandler.Wrapper(
		apphandler.NewContext[model.ListCityByIDRequest[[]int]](ctx, c.Mapper),
		func(ctx *apphandler.Context[model.ListCityByIDRequest[[]int], entity.City, model.CityResponse]) (any, int64, error) {
			return c.UseCase.List(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *City) GetById(ctx *fiber.Ctx) error {
	return apphandler.Wrapper(
		apphandler.NewContext[model.GetCityByIDRequest[[]int]](ctx, c.Mapper),
		func(ctx *apphandler.Context[model.GetCityByIDRequest[[]int], entity.City, model.CityResponse]) (any, int64, error) {
			return c.UseCase.GetById(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *City) GetFirstById(ctx *fiber.Ctx) error {
	return apphandler.Wrapper(
		apphandler.NewContext[model.GetCityByIDRequest[int]](ctx, c.Mapper),
		func(ctx *apphandler.Context[model.GetCityByIDRequest[int], entity.City, model.CityResponse]) (any, int64, error) {
			return c.UseCase.GetFirstById(ctx.Ctx, ctx.Request)
		},
	)
}
