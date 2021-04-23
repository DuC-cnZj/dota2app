package contracts

import (
	"net/http"
	"os"

	"github.com/DuC-cnZj/dota2app/pkg/config"
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

	FileManager() StorageInterface
	SetFileManager(StorageInterface)

	Run() chan os.Signal
	Shutdown()

	RegisterBeforeShutdownFunc(ShutdownFunc)
	RegisterAfterShutdownFunc(ShutdownFunc)

	EventDispatcher() DispatcherInterface
	SetEventDispatcher(DispatcherInterface)

	HttpHandler() http.Handler
	SetHttpHandler(http.Handler)
}
