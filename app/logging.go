package app

import (
	"api-ecommerce/config"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

func InitLogger() {
	// l.SetFlags(0)
	// l.SetOutput(new(logWriter))

	Formatter := new(log.TextFormatter)
	Formatter.ForceColors = true

	nestedFormatter := &nested.Formatter{
		HideKeys:        true,
		ShowFullLevel:   false,
		CallerFirst:     false,
		TimestampFormat: "Jan _2 15:04:05",
		CustomCallerFormatter: func(f *runtime.Frame) string {
			splits := strings.Split(f.File, "/")
			return fmt.Sprintf("[%s - %s:%d]", splits[len(splits)-1], path.Base(f.Function), f.Line)
		},
		FieldsOrder: []string{"component", "category", "method", "context", "params"},
	}

	switch {
	case config.LoadENV().ENVIRONTMENT == config.ENVIRONTMENT_PRODUCTION || config.LoadENV().ENVIRONTMENT == config.ENVIRONTMENT_TESTING:
		log.SetReportCaller(true)
		log.SetLevel(log.InfoLevel)
		log.SetFormatter(LocalFormatter{nestedFormatter})

	case strings.EqualFold(config.LoadENV().ENVIRONTMENT, config.ENVIRONTMENT_STAGING):
		log.SetReportCaller(true)
		log.SetLevel(log.DebugLevel)
		log.SetFormatter(LocalFormatter{nestedFormatter})

	case config.LoadENV().ENVIRONTMENT == "development" || config.LoadENV().ENVIRONTMENT == "staging_development" || strings.EqualFold(config.LoadENV().ENVIRONTMENT, config.ENVIRONTMENT_LOCAL):
		log.SetReportCaller(true)
		log.SetLevel(log.DebugLevel)
		log.SetFormatter(LocalFormatter{nestedFormatter})
		log.Info("kesini ?")
	}
}

type LocalFormatter struct {
	log.Formatter
}

func (u LocalFormatter) Format(e *log.Entry) ([]byte, error) {
	e.Time = e.Time.Add(time.Hour * 7)
	// e.Time = e.Time.Local()
	return u.Formatter.Format(e)
}

var reset = "\033[0m"
var purple = "\033[35m"

func EventLog(username string, event string) {
	log.Info(purple + "EVENT: " + reset + username + " " + event)
}
