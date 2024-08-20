package config

import (
	"github.com/aikuci/go-subdivisions-id/internal/delivery/http"
	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/route"
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
	provinceRepository := repository.NewProvinceRepository(config.Log)

	// setup use cases
	provinceUseCase := usecase.NewProvinceUseCase(config.Log, config.DB, provinceRepository)

	// setup controllers
	provinceController := http.NewProvinceController(config.Log, provinceUseCase)

	routeConfig := route.RouteConfig{
		App:                config.App,
		DB:                 config.DB,
		ProvinceController: provinceController,
	}
	routeConfig.Setup()
}
