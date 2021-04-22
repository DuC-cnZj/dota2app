package utils

import (
	"github.com/DuC-cnZj/dota2app/pkg/app/instance"
	"github.com/DuC-cnZj/dota2app/pkg/config"
	"github.com/DuC-cnZj/dota2app/pkg/contracts"
	"gorm.io/gorm"
)

func App() contracts.ApplicationInterface {
	return instance.App()
}

func Config() *config.Config {
	return App().Config()
}

func DB() *gorm.DB {
	return App().DBManager().DB()
}

func Event() contracts.DispatcherInterface {
	return App().EventDispatcher()
}
