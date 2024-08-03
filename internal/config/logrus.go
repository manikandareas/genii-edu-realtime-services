package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// NewLogger creates a new instance of logrus.Logger with the provided viper configuration.
// It sets the log level and formatter based on the configuration.
func NewLogger(viper *viper.Viper) *logrus.Logger {
	log := logrus.New()

	log.SetLevel(logrus.Level(viper.GetInt32("log.level")))
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}
