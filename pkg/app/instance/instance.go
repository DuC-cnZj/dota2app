package instance

import (
	"sync"

	"github.com/DuC-cnZj/dota2app/pkg/contracts"
)

var (
	app  contracts.ApplicationInterface
	once sync.Once
)

func SetInstance(app contracts.ApplicationInterface) {
	once.Do(func() {
		app = app
	})
}

func App() contracts.ApplicationInterface {
	return app
}
