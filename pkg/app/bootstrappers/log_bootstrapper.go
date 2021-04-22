package bootstrappers

import (
	"errors"
	"github.com/DuC-cnZj/dota2app/pkg/contracts"
	"github.com/DuC-cnZj/dota2app/pkg/dlog"
	log "github.com/sirupsen/logrus"

	"os"
)

type LogBootstrapper struct{}

func (a *LogBootstrapper) Bootstrap(app contracts.ApplicationInterface) error {
	switch app.Config().LogChannel {
	case "logrus":
		dlog.SetLogger(logrusLogger(app))
	default:
		return errors.New("log channel not exists: " + app.Config().LogChannel)
	}
	dlog.Debug("LogBootstrapper booted!")

	return nil
}

func logrusLogger(app contracts.ApplicationInterface) contracts.LoggerInterface {
	logger := log.New()

	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006/01/02 15:04:05",
	})

	if app.IsDebug() {
		logger.SetLevel(log.DebugLevel)
	} else {
		logger.SetFormatter(&log.JSONFormatter{})
		logger.SetLevel(log.InfoLevel)
	}

	return logger
}
