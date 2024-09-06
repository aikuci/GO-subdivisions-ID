package test

import (
	"os"

	"github.com/aikuci/go-subdivisions-id/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var app *fiber.App

var db *gorm.DB

var viperConfig *viper.Viper

var log *zap.Logger

func init() {
	viperConfig = config.NewViper()
	log = config.NewZapLog(viperConfig)
	app = config.NewFiber(viperConfig, &config.AppOptions{LogWriter: os.Stdout})
	db = config.NewDatabase(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		App:    app,
		Config: viperConfig,
		Log:    log,
		DB:     db,
	})
}
