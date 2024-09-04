package config

import (
	"fmt"
	"io"
	"log"
	"runtime/debug"
	"time"

	apperror "github.com/aikuci/go-subdivisions-id/pkg/util/error"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/storage/sqlite3/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type AppOptions struct {
	LogWriter io.Writer
}

func NewFiber(config *viper.Viper, options *AppOptions) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:                  config.GetString("app.name"),
		ErrorHandler:             NewErrorHandler(config),
		Prefork:                  config.GetBool("web.prefork"),
		EnableSplittingOnParsers: true,
	})

	app.Use(requestid.New(), logger.New(logger.Config{
		Format:     "[${time}](${pid} ${locals:requestid}) ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: time.RFC1123Z,
		Output:     options.LogWriter,
	}))
	app.Use(cache.New(cache.Config{
		KeyGenerator: func(c *fiber.Ctx) string {
			return utils.CopyString(c.Path()) + utils.CopyString(string(c.Request().URI().QueryString()))
		},
	}))
	app.Use(newLimiterConfig(config))
	app.Use(recover.New(recover.Config{
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			fmt.Println(c.Request().URI())
			stacks := fmt.Sprintf("panic: %v\n%s\n", e, debug.Stack())
			log.Println(stacks)
		},
		EnableStackTrace: true,
	}))

	return app
}

func NewErrorHandler(config *viper.Viper) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
		if e, ok := err.(*apperror.CustomErrorResponse); ok {
			code = e.HTTPCode
		}

		zapLog := NewZapLog(config)
		if rid, ok := ctx.Locals("requestid").(string); ok {
			zapLog = zapLog.With(zap.String("requestid", rid))
		}
		zapLog.Warn(err.Error())

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}

func newLimiterConfig(config *viper.Viper) fiber.Handler {
	storage := sqlite3.New(sqlite3.Config{
		Database:        "./storage/log/fiber-limiter.sqlite3",
		Table:           "fiber_storage",
		Reset:           false,
		GCInterval:      10 * time.Second,
		MaxOpenConns:    100,
		MaxIdleConns:    100,
		ConnMaxLifetime: 1 * time.Second,
	})

	limiterConfig := func() limiter.Config {
		if config.GetString("app.mode") == "production" {
			return newProductionLimiterConfig(storage)
		}

		return newDevelopmentLimiterConfig(storage)
	}()

	return limiter.New(limiterConfig)
}

func newDevelopmentLimiterConfig(storage fiber.Storage) limiter.Config {
	return limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Storage: storage,
	}
}

func newProductionLimiterConfig(storage fiber.Storage) limiter.Config {
	return limiter.Config{
		Max:                    10,
		SkipSuccessfulRequests: true,
		Storage:                storage,
	}
}
