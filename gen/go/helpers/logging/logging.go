package logging

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// GetZapLogger returns a zap logger configured based on the provided viper configuration.
// If the "KUBERNETES_SERVICE_HOST" environment variable is set, it uses the production configuration;
// otherwise, it uses the development configuration. The log level is set based on the "LOG_LEVEL" configuration value.
// If the log level is not set or invalid, it defaults to the zap logger's default level.
func NewZapLogger(v *viper.Viper) *zap.Logger {
	var zap_logger *zap.Logger
	if v.IsSet("KUBERNETES_SERVICE_HOST") {
		zap_config := zap.NewProductionConfig()
		if lvl, err := zap.ParseAtomicLevel(v.GetString("LOG_LEVEL")); err == nil {
			zap_config.Level = lvl
		}
		zap_logger, _ = zap_config.Build()
	} else {
		zap_config := zap.NewDevelopmentConfig()
		if lvl, err := zap.ParseAtomicLevel(v.GetString("LOG_LEVEL")); err == nil {
			zap_config.Level = lvl
		}
		zap_logger, _ = zap_config.Build()
	}
	return zap_logger
}
