package route

import (
	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"gorm.io/gorm"
)

type RouteConfig struct {
	App             *fiber.App
	DB              *gorm.DB
	ProvinceHandler *handler.Province
	CityHandler     *handler.City
	DistrictHandler *handler.District
	VillageHandler  *handler.Village
}

func (c *RouteConfig) Setup() {
	c.SetupRootRoute()
	c.SetupV1Route()
}

func (c *RouteConfig) SetupRootRoute() {
	c.App.Use(healthcheck.New())

	c.App.Get("/metrics", monitor.New())
	c.App.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("PONG")
	})
}
