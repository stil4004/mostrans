package ai_client

import (
	"context"
	"service/config"

	"github.com/go-resty/resty/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.uber.org/zap"
)

type ClientAIHttp struct {
	httpClient1 *resty.Client
	// httpClientEu *resty.Client
	// httpClientRu *resty.Client
}

func New(
	cfg *config.Config,

) (UseCase, error) {
	return &ClientAIHttp{
		httpClient1: resty.New().
			SetBaseURL("http://80.93.190.50:5000"),
		// httpClientEu: resty.New().
		// 	SetBaseURL(cfg.AiClient.VendorEUAI),
		// httpClientRu: resty.New().
		// 	SetBaseURL(cfg.AiClient.VendorRUAI),
	}, nil
}

type ZapLoggerAdapter struct {
	logger *zap.Logger
}

func (a *ZapLoggerAdapter) Log(ctx context.Context, level logging.Level, msg string, fields ...interface{}) {
	switch level {
	case logging.LevelDebug:
		a.logger.Debug(msg, convertFields(fields)...)
	case logging.LevelInfo:
		a.logger.Info(msg, convertFields(fields)...)
	case logging.LevelWarn:
		a.logger.Warn(msg, convertFields(fields)...)
	case logging.LevelError:
		a.logger.Error(msg, convertFields(fields)...)
	}
}

func convertFields(fields []interface{}) []zap.Field {
	var zapFields []zap.Field
	for i := 0; i < len(fields); i += 2 {
		if i+1 < len(fields) {
			zapFields = append(zapFields, zap.Any(fields[i].(string), fields[i+1]))
		}
	}
	return zapFields
}

func InterceptorLogger(logger *zap.Logger) logging.Logger {
	return &ZapLoggerAdapter{logger: logger}
}
