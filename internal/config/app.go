package config

import (
	"github.com/aikuci/go-subdivisions-id/internal/delivery/http"
	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/route"
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/repository"
	"github.com/aikuci/go-subdivisions-id/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	App      *fiber.App
	Config   *viper.Viper
	Log      *zap.Logger
	Validate *validator.Validate
	DB       *gorm.DB
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	crudRepository := repository.NewCrudRepository[entity.Province](config.Log)
	provinceRepository := repository.NewProvinceRepository(config.Log)

	// setup use cases
	crudUseCase := usecase.NewCrudUseCase(config.Log, config.DB, crudRepository)
	provinceUseCase := usecase.NewProvinceUseCase(config.Log, config.DB, provinceRepository)

	// setup controllers
	crudController := http.NewCrudController[model.ProvinceResponse](config.Log, crudUseCase)
	provinceController := http.NewProvinceController(config.Log, provinceUseCase)

	routeConfig := route.RouteConfig{
		App:                config.App,
		DB:                 config.DB,
		CrudController:     crudController,
		ProvinceController: provinceController,
	}
	routeConfig.Setup()
}
