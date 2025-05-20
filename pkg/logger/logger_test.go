package logger_test

import (
	"context"
	"github.com/mhthrh/common_pkg/pkg/logger"
	"go.uber.org/zap"
	"testing"
)

var (
	l logger.ILogger
)

func init() {
	l = logger.NewLogger("test")
}

func TestNewLogger(t *testing.T) {
	_ = logger.NewLogger("test")
}
func TestLog_Info(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, logger.UserIDContext, "value-1")
	l.Info(ctx, "key", zap.String("key1", "value1"))
}
func TestLog_Debug(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, logger.SpanIDContext, "value-2")
	l.Info(ctx, "key", zap.String("key1", "value1"))
}
func TestLog_Error(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, logger.TransactionIDContext, "value-3")
	l.Info(ctx, "key", zap.String("key1", "value1"))
}
func TestLog_Warn(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, logger.TraceIDContext, "value-4")
	l.Info(ctx, "key", zap.String("key1", "value1"))
}
