package http

import (
	"context"

	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/middleware/requestid"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Controller struct {
	context context.Context
	log     *zap.Logger
}

func Prepare(ctx *fiber.Ctx, log *zap.Logger) *Controller {
	context := requestid.SetContext(ctx.UserContext(), ctx)

	return &Controller{
		context: context,
		log:     log.With(zap.String("requestid", requestid.FromContext(context))),
	}
}
