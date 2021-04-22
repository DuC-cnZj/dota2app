package bootstrappers

import (
	"github.com/DuC-cnZj/dota2app/pkg/contracts"
	t "github.com/DuC-cnZj/dota2app/pkg/translator"
)

type I18nBootstrapper struct{}

func (i *I18nBootstrapper) Bootstrap(app contracts.ApplicationInterface) error {
	t.Init()

	return nil
}
