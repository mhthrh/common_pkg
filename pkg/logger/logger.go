package logger

import (
	"context"
	"go.elastic.co/apm/module/apmzap"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Log struct {
	log *zap.Logger
}

func NewLogger() ILogger {
	encoderConfig := ecszap.ECSCompatibleEncoderConfig(zap.NewProductionEncoderConfig())
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	logger := zap.New(zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), zap.NewAtomicLevel()), zap.WrapCore((&apmzap.Core{}).WrapCore), zap.AddCaller())
	logger = logger.Named("com.my-app.user")
	zap.ReplaceGlobals(logger)
	return Log{log: logger}
}
func (l Log) Info(ctx context.Context, msg string, args ...zap.Field) {
	l.log.With(zap.Field{
		Key:       "transactionInfo",
		Type:      zapcore.StringType,
		Integer:   0,
		String:    "",
		Interface: getLoggerFieldsFromCtx(ctx),
	}).Info(buildMsgField(msg, args...))

}

func (l Log) Debug(ctx context.Context, msg string, args ...zap.Field) {

	l.log.With(zap.Field{
		Key:       "transactionInfo",
		Type:      zapcore.StringType,
		Integer:   0,
		String:    "",
		Interface: getLoggerFieldsFromCtx(ctx),
	}).Debug(buildMsgField(msg, args...))
}

func (l Log) Warn(ctx context.Context, msg string, args ...zap.Field) {
	l.log.With(zap.Field{
		Key:       "transactionInfo",
		Type:      zapcore.StringType,
		Integer:   0,
		String:    "",
		Interface: getLoggerFieldsFromCtx(ctx),
	}).Warn(buildMsgField(msg, args...))

}

func (l Log) Error(ctx context.Context, msg string, args ...zap.Field) {
	l.log.With(zap.Field{
		Key:       "transactionInfo",
		Type:      zapcore.StringType,
		Integer:   0,
		String:    "",
		Interface: getLoggerFieldsFromCtx(ctx),
	}).Error(buildMsgField(msg, args...))

}

func (l Log) Fatal(ctx context.Context, msg string, args ...zap.Field) {
	l.log.With(zap.Field{
		Key:       "transactionInfo",
		Type:      zapcore.StringType,
		Integer:   0,
		String:    "",
		Interface: getLoggerFieldsFromCtx(ctx),
	}).Fatal(buildMsgField(msg, args...))
}
