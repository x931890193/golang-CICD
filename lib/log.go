package lib

import (
	"os"
	"github.com/sirupsen/logrus"
)

var Logs = logrus.New()

func init() {
	Logs.Out = os.Stdout
	Logs.Formatter = &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05.000"}
}
