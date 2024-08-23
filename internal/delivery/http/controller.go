package http

import (
	"context"

	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/middleware/requestid"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Controller struct {
	context context.Context
	Log     *zap.Logger
}

func NewController(log *zap.Logger) *Controller {
	return &Controller{
		Log: log,
	}
}

func (c *Controller) Prepare(ctx *fiber.Ctx) {
	c.context = requestid.SetContext(ctx.UserContext(), ctx)
	c.Log = c.Log.With(zap.String("requestid", requestid.FromContext(c.context)))
}
