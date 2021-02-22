package option_a

import (
	"fmt"
	"os"
	"strings"

	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/config"
	"github.com/sirupsen/logrus"
)

var (
	//Log ...
	Log *logrus.Logger
)

func init() {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		level = logrus.DebugLevel
	}
	Log = &logrus.Logger{
		Level:     level,
		Out:       os.Stdout,
		Formatter: &logrus.JSONFormatter{},
	}
}

//Debug ...
func Debug(msg string, tags ...string) {
	if Log.Level < logrus.DebugLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Debug(msg)
}

//Info ..
func Info(msg string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Info(msg)
}

//Error ...
func Error(msg string, err error, tags ...string) {
	if Log.Level < logrus.ErrorLevel {
		return
	}
	msg = fmt.Sprintf("%s - ERROR - %s", msg, err)
	Log.WithFields(parseFields(tags...)).Error(msg)
}

func parseFields(tags ...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))
	for _, tag := range tags {
		elements := strings.Split(tag, ":")
		result[strings.TrimSpace(elements[0])] = strings.TrimSpace(elements[1])
	}
	return result
}
