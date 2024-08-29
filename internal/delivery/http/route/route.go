package route

import (
	"github.com/aikuci/go-subdivisions-id/internal/delivery/http"
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	apphttp "github.com/aikuci/go-subdivisions-id/pkg/delivery/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"gorm.io/gorm"
)

type RouteConfig struct {
	App                *fiber.App
	DB                 *gorm.DB
	ProvinceController *apphttp.CrudController[entity.Province, model.ProvinceResponse]
	CityController     *http.CityController
	DistrictController *http.DistrictController
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
