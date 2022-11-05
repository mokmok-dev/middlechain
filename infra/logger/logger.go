package logger

import (
	"fmt"

	domain "github.com/Pranc1ngPegasus/golang-template/domain/logger"
	"github.com/google/wire"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

var _ domain.Logger = (*Logger)(nil)

type Logger struct {
	logger *zap.Logger
}

var NewLoggerSet = wire.NewSet(
	wire.Bind(new(domain.Logger), new(*Logger)),
	NewLogger,
)

func NewLogger() (*Logger, error) {
	log, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	return &Logger{
		logger: log,
	}, nil
}

func (l *Logger) Field(key string, message interface{}) domain.Field {
	return domain.Field{
		Key:       key,
		Interface: message,
	}
}

func (l *Logger) field(field domain.Field) zap.Field {
	switch i := field.Interface.(type) {
	case string:
		return zap.String(field.Key, i)
	case int:
		return zap.Int(field.Key, i)
	case bool:
		return zap.Bool(field.Key, i)
	default:
		return zap.Any(field.Key, i)
	}
}

func (l *Logger) Info(message string, fields ...domain.Field) {
	zapfields := lo.Map(fields, func(field domain.Field, _ int) zap.Field {
		return l.field(field)
	})

	l.logger.Info(message, zapfields...)
}

func (l *Logger) Error(message string, err error, fields ...domain.Field) {
	zapfields := lo.Map(fields, func(field domain.Field, _ int) zap.Field {
		return l.field(field)
	})

	zapfields = append(zapfields, zap.Error(err))

	l.logger.Error(message, zapfields...)
}
