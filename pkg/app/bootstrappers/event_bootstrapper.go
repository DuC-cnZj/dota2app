package bootstrappers

import (
	"github.com/DuC-cnZj/dota2app/pkg/contracts"
	"github.com/DuC-cnZj/dota2app/pkg/dlog"
	"github.com/DuC-cnZj/dota2app/pkg/event"
)

type EventBootstrapper struct{}

var events map[contracts.Event][]contracts.Listener = map[contracts.Event][]contracts.Listener{}

func (e *EventBootstrapper) Bootstrap(app contracts.ApplicationInterface) error {
	app.SetEventDispatcher(event.NewDispatcher(app))

	for e, listeners := range events {
		for _, listener := range listeners {
			app.EventDispatcher().Listen(e, listener)
		}
	}

	dlog.Debug("EventBootstrapper booted.")

	return nil
}
