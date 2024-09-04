package config

import (
	"fmt"
	"io"
	"log"
	"runtime/debug"
	"time"

	apperror "github.com/aikuci/go-subdivisions-id/pkg/util/error"
	applog "github.com/aikuci/go-subdivisions-id/pkg/util/log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/storage/sqlite3/v2"
	"github.com/spf13/viper"
)

type AppOptions struct {
	LogWriter io.Writer
}

func NewFiber(viper *viper.Viper, options *AppOptions) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:                  viper.GetString("app.name"),
		ErrorHandler:             NewErrorHandler(viper),
		Prefork:                  viper.GetBool("web.prefork"),
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
	app.Use(newLimiterConfig(viper))
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

func NewErrorHandler(viper *viper.Viper) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		switch e := err.(type) {
		case *fiber.Error:
			code = e.Code
		case *apperror.CustomErrorResponse:
			code = e.HTTPCode
		}

		applog.Write(NewZapLog(viper), ctx.UserContext(), "[fiber]: ", err)

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}

func newLimiterConfig(viper *viper.Viper) fiber.Handler {
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
		if viper.GetString("app.mode") == "production" {
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
