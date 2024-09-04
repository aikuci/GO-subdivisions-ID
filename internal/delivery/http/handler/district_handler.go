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

type DistrictHandler struct {
	Log     *zap.Logger
	UseCase usecase.DistrictUseCase
	Mapper  appmapper.CruderMapper[entity.District, model.DistrictResponse]
}

func NewDistrictHandler(log *zap.Logger, useCase *usecase.DistrictUseCase, mapper appmapper.CruderMapper[entity.District, model.DistrictResponse]) *DistrictHandler {
	return &DistrictHandler{
		Log:     log,
		UseCase: *useCase,
		Mapper:  mapper,
	}
}

func (c *DistrictHandler) List(ctx *fiber.Ctx) error {
	return apphandler.Wrapper(
		apphandler.NewContext[model.ListDistrictByIDRequest[[]int]](c.Log, ctx, c.Mapper),
		func(ctx *apphandler.HandlerContext[model.ListDistrictByIDRequest[[]int], entity.District, model.DistrictResponse]) (any, int64, error) {
			return c.UseCase.List(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *DistrictHandler) GetById(ctx *fiber.Ctx) error {
	return apphandler.Wrapper(
		apphandler.NewContext[model.GetDistrictByIDRequest[[]int]](c.Log, ctx, c.Mapper),
		func(ctx *apphandler.HandlerContext[model.GetDistrictByIDRequest[[]int], entity.District, model.DistrictResponse]) (any, int64, error) {
			return c.UseCase.GetById(ctx.Ctx, ctx.Request)
		},
	)
}

func (c *DistrictHandler) GetFirstById(ctx *fiber.Ctx) error {
	return apphandler.Wrapper(
		apphandler.NewContext[model.GetDistrictByIDRequest[int]](c.Log, ctx, c.Mapper),
		func(ctx *apphandler.HandlerContext[model.GetDistrictByIDRequest[int], entity.District, model.DistrictResponse]) (any, int64, error) {
			return c.UseCase.GetFirstById(ctx.Ctx, ctx.Request)
		},
	)
}
