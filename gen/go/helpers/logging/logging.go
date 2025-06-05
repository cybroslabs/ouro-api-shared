package logging

import (
	"os"

	"go.uber.org/zap"
)

// NewZapLogger returns a zap logger configured based on the provided parameters.
// It checks if the application is running in a Kubernetes environment by looking for the KUBERNETES_SERVICE_HOST environment variable.
// If it is running in Kubernetes, it uses the production configuration; otherwise, it uses the development configuration.
// The logger is configured to use the log level specified by the LOG_LEVEL environment variable.
// If the LOG_LEVEL is not set or is invalid, it defaults to the zap logger's default level.
func NewZapLogger() *zap.Logger {
	var zap_logger *zap.Logger
	var logLevel = os.Getenv("LOG_LEVEL")
	if _, ok := os.LookupEnv("KUBERNETES_SERVICE_HOST"); ok {
		zap_config := zap.NewProductionConfig()
		if lvl, err := zap.ParseAtomicLevel(logLevel); err == nil {
			zap_config.Level = lvl
		}
		zap_logger, _ = zap_config.Build()
	} else {
		zap_config := zap.NewDevelopmentConfig()
		if lvl, err := zap.ParseAtomicLevel(logLevel); err == nil {
			zap_config.Level = lvl
		}
		zap_logger, _ = zap_config.Build()
	}
	return zap_logger
}
