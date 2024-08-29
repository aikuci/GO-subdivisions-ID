package config

import (
	"github.com/aikuci/go-subdivisions-id/internal/delivery/http"
	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/route"
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model/mapper"
	"github.com/aikuci/go-subdivisions-id/internal/repository"
	"github.com/aikuci/go-subdivisions-id/internal/usecase"
	apphttp "github.com/aikuci/go-subdivisions-id/pkg/delivery/http"
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
	// setup repositories
	provinceRepository := apprepository.NewCrudRepository[entity.Province, int, []int]()
	cityRepository := repository.NewCityRepository[int, []int]()
	districtRepository := repository.NewDistrictRepository[int, []int]()

	// setup use cases
	provinceUseCase := appusecase.NewCrudUseCase(config.Log, config.DB, provinceRepository)
	cityUseCase := usecase.NewCityUseCase(config.Log, config.DB, cityRepository)
	districtUseCase := usecase.NewDistrictUseCase(config.Log, config.DB, districtRepository)

	// setup controllers
	provinceController := apphttp.NewCrudController(config.Log, provinceUseCase, mapper.NewProvinceMapper())
	cityController := http.NewCityController(config.Log, cityUseCase, mapper.NewCityMapper())
	districtController := http.NewDistrictController(config.Log, districtUseCase, mapper.NewDistrictMapper())

	routeConfig := route.RouteConfig{
		App:                config.App,
		DB:                 config.DB,
		ProvinceController: provinceController,
		CityController:     cityController,
		DistrictController: districtController,
	}
	routeConfig.Setup()
}
