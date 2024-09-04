package config

import (
	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/handler"
	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/route"
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/repository"
	"github.com/aikuci/go-subdivisions-id/internal/usecase"
	apprepository "github.com/aikuci/go-subdivisions-id/pkg/repository"
	appusecase "github.com/aikuci/go-subdivisions-id/pkg/usecase"

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
	// Repositories
	provinceRepository := apprepository.NewCrud[entity.Province, int, []int]()
	cityRepository := repository.NewCity[int, []int]()
	districtRepository := repository.NewDistrict[int, []int]()
	villageRepository := repository.NewVillage[int, []int]()

	// UseCases
	provinceUseCase := appusecase.NewCrud(config.Log, config.DB, provinceRepository)
	cityUseCase := usecase.NewCity(config.Log, config.DB, cityRepository)
	districtUseCase := usecase.NewDistrict(config.Log, config.DB, districtRepository)
	villageUseCase := usecase.NewVillage(config.Log, config.DB, villageRepository)

	// Handlers
	provinceHandler := handler.NewProvince(provinceUseCase)
	cityHandler := handler.NewCity(cityUseCase)
	districtHandler := handler.NewDistrict(districtUseCase)
	villageHandler := handler.NewVillage(villageUseCase)

	routeConfig := route.RouteConfig{
		App:             config.App,
		DB:              config.DB,
		ProvinceHandler: provinceHandler,
		CityHandler:     cityHandler,
		DistrictHandler: districtHandler,
		VillageHandler:  villageHandler,
	}
	routeConfig.Setup()
}
