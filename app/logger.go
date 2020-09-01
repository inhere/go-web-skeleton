package app

import (
	"github.com/sirupsen/logrus"
)

const RFC3339Normal = "2006-01-02 15:04:05"
const RFC3339Shorted = "2006/01/02 15:04:05"
const RFC3339NanoFixed = "2006-01-02T15:04:05.000000000Z07:00"

// initLog init log setting
func InitLogger() error {
	// newGenericLogger()
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: RFC3339Normal,
		DisableColors:   false,
		FullTimestamp:   true,
	})

	logrus.Info("logger construction succeeded")

	return nil
}
