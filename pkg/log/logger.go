package log

import (
	"go.uber.org/zap/zapcore"
	config "go-clean-architecture/config"
	"strings"
	"time"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	logLevel := config.GetConfiguration().LogLevel
	runEnv := config.GetConfiguration().RunEnv
	var configLogger zap.Config
	if strings.ToUpper(runEnv) == "PROD" {
		configLogger = zap.NewProductionConfig()
	} else {
		configLogger = zap.NewDevelopmentConfig()
	}
	configLogger.EncoderConfig.EncodeTime = syslogTimeEncoder
	configLogger.Level.UnmarshalText([]byte(logLevel))
	log, _ := configLogger.Build()
	logger = log.Named("service").Sugar()
}

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func GetLogger() *zap.SugaredLogger {
	return logger
}
