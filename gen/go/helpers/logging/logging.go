package logging

import (
	"go.uber.org/zap"
)

// GetZapLogger returns a zap logger configured based on the provided parameters.
// If inKubernetes is true, it uses production configuration; otherwise, it uses development configuration.
// If logLevel is provided, it sets the logger's level accordingly.
// The function returns a pointer to a zap.Logger instance.
func NewZapLogger(inKubernetes bool, logLevel *string) *zap.Logger {
	var zap_logger *zap.Logger
	if inKubernetes {
		zap_config := zap.NewProductionConfig()
		if logLevel != nil {
			if lvl, err := zap.ParseAtomicLevel(*logLevel); err == nil {
				zap_config.Level = lvl
			}
		}
		zap_logger, _ = zap_config.Build()
	} else {
		zap_config := zap.NewDevelopmentConfig()
		if logLevel != nil {
			if lvl, err := zap.ParseAtomicLevel(*logLevel); err == nil {
				zap_config.Level = lvl
			}
		}
		zap_logger, _ = zap_config.Build()
	}
	return zap_logger
}
