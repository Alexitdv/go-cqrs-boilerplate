package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// InitLogger inits logger
func InitLogger(logLevel string) error {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return fmt.Errorf("failed to parse log level. %v", err)
	}
	logrus.SetLevel(lvl)

	return nil
}
