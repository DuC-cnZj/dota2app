package bootstrappers

import (
	"github.com/DuC-cnZj/dota2app/pkg/contracts"
	"github.com/DuC-cnZj/dota2app/pkg/uploader"
)

type StorageBootstrapper struct{}

func (f *StorageBootstrapper) Bootstrap(app contracts.ApplicationInterface) error {
	m, e := uploader.Init(app)
	if e != nil {
		return e
	}

	app.SetFileManager(m)
	return nil
}
