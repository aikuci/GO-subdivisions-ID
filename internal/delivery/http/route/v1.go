package route

import "github.com/gofiber/fiber/v2"

func (c *RouteConfig) SetupV1Route() {
	v1 := c.App.Group("/v1")

	v1.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("PONG from V1")
	})

	v1.Get("/provinces", c.ProvinceHandler.CrudHandler.List)
	v1.Get("/provinces/:id<int>", c.ProvinceHandler.CrudHandler.GetFirstById)

	v1.Get("/cities", c.CityHandler.List)
	v1.Get("/cities/:id<int>", c.CityHandler.GetById)
	v1.Get("/provinces/:id_province<int>/cities", c.CityHandler.GetById)
	v1.Get("/provinces/:id_province<int>/cities/:id<int>", c.CityHandler.GetFirstById)

	v1.Get("/districts", c.DistrictHandler.List)
	v1.Get("/districts/:id<int>", c.DistrictHandler.GetById)
	v1.Get("/provinces/:id_province<int>/cities/:id_city<int>/districts", c.DistrictHandler.GetById)
	v1.Get("/provinces/:id_province<int>/cities/:id_city<int>/districts/:id<int>", c.DistrictHandler.GetFirstById)

	v1.Get("/villages", c.VillageHandler.List)
	v1.Get("/villages/:id<int>", c.VillageHandler.GetById)
	v1.Get("/provinces/:id_province<int>/cities/:id_city<int>/districts/:id_district<int>/villages", c.VillageHandler.GetById)
	v1.Get("/provinces/:id_province<int>/cities/:id_city<int>/districts/:id_district<int>/villages/:id<int>", c.VillageHandler.GetFirstById)
}
