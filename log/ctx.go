package log

import (
	"context"
)

var _ctxKeyLogger = struct{}{}

func LoggerWithContext(ctx context.Context, l *Logger) context.Context {
	return context.WithValue(ctx, _ctxKeyLogger, l)
}

func LoggerFromContext(ctx context.Context) *Logger {
	return ctx.Value(_ctxKeyLogger).(*Logger)
}
