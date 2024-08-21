package route

import (
	"github.com/aikuci/go-subdivisions-id/internal/delivery/http"
	"github.com/aikuci/go-subdivisions-id/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"gorm.io/gorm"
)

type RouteConfig struct {
	App                *fiber.App
	DB                 *gorm.DB
	ProvinceController *http.CrudController[model.ProvinceResponse]
	CityController     *http.CrudController[model.CityResponse]
}

func (c *RouteConfig) Setup() {
	c.SetupRootRoute()
	c.SetupV1Route()
}

func (c *RouteConfig) SetupRootRoute() {
	c.App.Use(healthcheck.New())

	c.App.Get("/metrics", monitor.New())
	c.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("PONG")
	})
}
