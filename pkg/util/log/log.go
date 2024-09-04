package log

import (
	"context"

	"github.com/aikuci/go-subdivisions-id/pkg/util/context/requestid"

	"go.uber.org/zap"
)

func Write(zapLog *zap.Logger, ctx context.Context, errorMessage string, err error) {
	zapLog = zapLog.WithOptions(zap.AddCallerSkip(1))

	if requestid.FromContext(ctx) != "" {
		zapLog = zapLog.With(zap.String("requestid", requestid.FromContext(ctx)))
	}

	zapLog.Warn(errorMessage, zap.Error(err))
}
