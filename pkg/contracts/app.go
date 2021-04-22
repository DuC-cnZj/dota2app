package contracts

import (
	"github.com/DuC-cnZj/dota2app/pkg/config"
	"net/http"
	"os"
)

type ShutdownFunc func(ApplicationInterface)

type Bootstrapper interface {
	Bootstrap(ApplicationInterface) error
}

type Option func(ApplicationInterface)

type ApplicationInterface interface {
	IsDebug() bool

	Bootstrap() error
	Config() *config.Config

	DBManager() DBManager

	Run() chan os.Signal
	Shutdown()

	RegisterBeforeShutdownFunc(ShutdownFunc)
	RegisterAfterShutdownFunc(ShutdownFunc)

	Event() DispatcherInterface
	SetEventDispatcher(DispatcherInterface)

	HttpHandler() http.Handler
	SetHttpHandler(http.Handler)
}
