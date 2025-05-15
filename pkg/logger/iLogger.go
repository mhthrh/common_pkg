package logger

import (
	"context"
	"go.uber.org/zap"
)

type ILogger interface {
	Info(context.Context, string, ...zap.Field)
	Debug(context.Context, string, ...zap.Field)
	Warn(context.Context, string, ...zap.Field)
	Error(context.Context, string, ...zap.Field)
	Fatal(context.Context, string, ...zap.Field)
}
