package route

import "github.com/gofiber/fiber/v2"

func (c *RouteConfig) SetupV1Route() {
	v1 := c.App.Group("/v1")

	v1.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("PONG from V1")
	})

	v1.Get("/provinces", c.ProvinceController.List)
	v1.Get("/provinces/:id<int>", c.ProvinceController.GetFirstById)

	v1.Get("/cities", c.CityController.List)
	v1.Get("/cities/:id<int>", c.CityController.GetById)
	v1.Get("/provinces/:id_province<int>/cities", c.CityController.GetById)
	v1.Get("/provinces/:id_province<int>/cities/:id<int>", c.CityController.GetFirstById)

	v1.Get("/districts", c.DistrictController.List)
	v1.Get("/districts/:id<int>", c.DistrictController.GetById)
	v1.Get("/provinces/:id_province<int>/cities/:id_city<int>/districts", c.DistrictController.GetById)
	v1.Get("/provinces/:id_province<int>/cities/:id_city<int>/districts/:id<int>", c.DistrictController.GetFirstById)

	v1.Get("/villages", c.VillageController.List)
	v1.Get("/villages/:id<int>", c.VillageController.GetById)
	v1.Get("/provinces/:id_province<int>/cities/:id_city<int>/districts/:id_district<int>/villages", c.VillageController.GetById)
	v1.Get("/provinces/:id_province<int>/cities/:id_city<int>/districts/:id_district<int>/villages/:id<int>", c.VillageController.GetFirstById)
}
