package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DuC-cnZj/dota2app/pkg/app/bootstrappers"
	"github.com/DuC-cnZj/dota2app/pkg/app/instance"
	"github.com/DuC-cnZj/dota2app/pkg/config"
	"github.com/DuC-cnZj/dota2app/pkg/contracts"
	"github.com/DuC-cnZj/dota2app/pkg/database"
	"github.com/DuC-cnZj/dota2app/pkg/dlog"
)

var _ contracts.ApplicationInterface = (*Application)(nil)

var DefaultBootstrappers = []contracts.Bootstrapper{
	&bootstrappers.I18nBootstrapper{},
	&bootstrappers.DBBootstrapper{},
	&bootstrappers.StorageBootstrapper{},
	&bootstrappers.WebBootstrapper{},
	&bootstrappers.RouterBootstrapper{},
}

type Application struct {
	config *config.Config

	bootstrappers []contracts.Bootstrapper

	storageManager contracts.StorageInterface
	dbManager      contracts.DBManager

	httpHandler http.Handler
	httpServer  *http.Server

	beforeShutdownFunctions []contracts.ShutdownFunc
	afterShutdownFunctions  []contracts.ShutdownFunc

	dispatcher contracts.DispatcherInterface
}

func (app *Application) FileManager() contracts.StorageInterface {
	return app.storageManager
}

func (app *Application) SetFileManager(m contracts.StorageInterface) {
	app.storageManager = m
}

func (app *Application) HttpHandler() http.Handler {
	return app.httpHandler
}

func (app *Application) SetHttpHandler(handler http.Handler) {
	app.httpHandler = handler
	app.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Config().AppPort),
		Handler: handler,
	}
}

func (app *Application) EventDispatcher() contracts.DispatcherInterface {
	return app.dispatcher
}

func (app *Application) SetEventDispatcher(dispatcher contracts.DispatcherInterface) {
	app.dispatcher = dispatcher
}

func NewApplication(config *config.Config, opts ...contracts.Option) *Application {
	var mustBooted = []contracts.Bootstrapper{
		&bootstrappers.LogBootstrapper{},
		&bootstrappers.EventBootstrapper{},
	}

	app := &Application{
		bootstrappers: DefaultBootstrappers,
		config:        config,
	}

	app.dbManager = database.NewManager(app)

	for _, opt := range opts {
		opt(app)
	}

	instance.SetInstance(app)

	for _, bootstrapper := range mustBooted {
		if err := bootstrapper.Bootstrap(app); err != nil {
			dlog.Fatal(err)
		}
	}

	return app
}

func (app *Application) Bootstrap() error {
	for _, bootstrapper := range app.bootstrappers {
		if err := bootstrapper.Bootstrap(app); err != nil {
			return err
		}
	}

	return nil
}

func (app *Application) Config() *config.Config {
	return app.config
}

func (app *Application) DBManager() contracts.DBManager {
	return app.dbManager
}

func (app *Application) IsDebug() bool {
	return app.config.Debug
}

func (app *Application) Run() chan os.Signal {
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	go func() {
		dlog.Infof("server running at %s.", app.httpServer.Addr)
		if err := app.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			dlog.Fatal(err)
		}
	}()

	if app.Config().ProfileEnabled {
		go func() {
			dlog.Info("Starting pprof server on localhost:6060.")
			if err := http.ListenAndServe("localhost:6060", nil); err != nil && err != http.ErrServerClosed {
				dlog.Error(err)
			}
		}()
	}

	return done
}

func (app *Application) Shutdown() {
	var err error

	callBeforeShutdownFunctions(app)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = app.httpServer.Shutdown(ctx)
	if err != nil {
		dlog.Error(err)
	}

	callAfterShutdownFunctions(app)

	dlog.Info("server graceful shutdown.")
}

func (app *Application) RegisterAfterShutdownFunc(fn contracts.ShutdownFunc) {
	app.beforeShutdownFunctions = append(app.afterShutdownFunctions, fn)
}

func (app *Application) RegisterBeforeShutdownFunc(fn contracts.ShutdownFunc) {
	app.beforeShutdownFunctions = append(app.beforeShutdownFunctions, fn)
}

func callAfterShutdownFunctions(app *Application) {
	for _, fn := range app.afterShutdownFunctions {
		fn(app)
	}
}

func callBeforeShutdownFunctions(app *Application) {
	for _, fn := range app.beforeShutdownFunctions {
		fn(app)
	}
}
