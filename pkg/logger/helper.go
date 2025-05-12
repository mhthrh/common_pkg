package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"math"
	"strings"
	"time"
)

const (
	TraceIDContext       = "TraceID"
	TransactionIDContext = "TransactionID"
	SpanIDContext        = "SpanID"
	UserIDContext        = "userID"
)

func getLoggerFieldsFromCtx(ctx context.Context) string {
	traceID := getTraceID(ctx)
	transactionID := getTransactionID(ctx)
	userID := getUserID(ctx)
	spanID := getSpanID(ctx)

	return fmt.Sprintf("trace.id: %s, transaction.id: %s, span.id: %s, user.id: %s", traceID, transactionID, spanID, userID)
}
func buildMsgField(msg string, fields ...zap.Field) string {
	Msgs := make([]string, len(fields))
	for _, field := range fields {
		if field.Type == zapcore.SkipType {
			continue
		}
		Msgs = append(Msgs, fmt.Sprintf("%s:%s", field.Key, field2String(field)))
	}
	return fmt.Sprintf("%s:%s", msg, strings.Join(Msgs, ";"))
}

func field2String(f zap.Field) string {
	switch f.Type {
	case zapcore.ArrayMarshalerType:
		return fmt.Sprintf("%v", f.Interface.(zapcore.ArrayMarshaler))
	case zapcore.ObjectMarshalerType:
		return fmt.Sprintf("%v", f.Interface.(zapcore.ObjectMarshaler))
	case zapcore.BinaryType:
		return fmt.Sprintf("%v", f.Interface.([]byte))
	case zapcore.BoolType:
		return fmt.Sprintf("%v", f.Integer == 1)
	case zapcore.ByteStringType:
		return fmt.Sprintf("%v", f.Interface.([]byte))
	case zapcore.Complex128Type:
		return fmt.Sprintf("%v", f.Interface.(complex128))
	case zapcore.Complex64Type:
		return fmt.Sprintf("%v", f.Interface.(complex64))
	case zapcore.DurationType:
		return fmt.Sprintf("%v", time.Duration(f.Integer))
	case zapcore.Float64Type:
		return fmt.Sprintf("%v", math.Float64frombits(uint64(f.Integer)))
	case zapcore.Float32Type:
		return fmt.Sprintf("%v", math.Float32frombits(uint32(f.Integer)))
	case zapcore.Int64Type:
		return fmt.Sprintf("%v", f.Integer)
	case zapcore.Int32Type:
		return fmt.Sprintf("%v", int32(f.Integer))
	case zapcore.Int16Type:
		return fmt.Sprintf("%v", int16(f.Integer))
	case zapcore.Int8Type:
		return fmt.Sprintf("%v", int8(f.Integer))
	case zapcore.StringType:
		return fmt.Sprintf("%v", f.String)
	case zapcore.TimeType:
		if f.Interface != nil {
			return fmt.Sprintf("%v", time.Unix(0, f.Integer).In(f.Interface.(*time.Location)))
		} else {
			return fmt.Sprintf("%v", time.Unix(0, f.Integer))
		}
	case zapcore.Uint64Type:
		return fmt.Sprintf("%v", uint64(f.Integer))
	case zapcore.Uint32Type:
		return fmt.Sprintf("%v", uint32(f.Integer))
	case zapcore.Uint16Type:
		return fmt.Sprintf("%v", uint16(f.Integer))
	case zapcore.Uint8Type:
		return fmt.Sprintf("%v", uint8(f.Integer))
	case zapcore.UintptrType:
		return fmt.Sprintf("%v", uintptr(f.Integer))
	case zapcore.ReflectType:
		return fmt.Sprintf("%+v", f.Interface)
	case zapcore.NamespaceType:
		return fmt.Sprintf("%v", f.Key)
	case zapcore.StringerType:
		return fmt.Sprintf("%v", f.Interface)
	case zapcore.ErrorType:
		return fmt.Sprintf("%v", f.Interface.(error).Error())
	case zapcore.SkipType:
		break
	default:
		return ""
	}
	return ""
}
func getTraceID(ctx context.Context) string {
	td := ctx.Value(TraceIDContext)
	if td != nil {
		return td.(string)
	}
	return ""
}

func getTransactionID(ctx context.Context) string {
	td := ctx.Value(TransactionIDContext)
	if td != nil {
		return td.(string)
	}
	return ""
}

func getSpanID(ctx context.Context) string {
	td := ctx.Value(SpanIDContext)
	if td != nil {
		return td.(string)
	}
	return ""
}

func getUserID(ctx context.Context) string {
	userId := ctx.Value(UserIDContext)
	if userId != nil {
		return userId.(string)
	}
	return ""
}
