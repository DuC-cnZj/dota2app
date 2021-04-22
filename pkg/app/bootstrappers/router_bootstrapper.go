package bootstrappers

import (
	"github.com/DuC-cnZj/dota2app/pkg/contracts"
	"github.com/DuC-cnZj/dota2app/pkg/router"
	"github.com/gin-gonic/gin"
)

type RouterBootstrapper struct{}

func (r *RouterBootstrapper) Bootstrap(app contracts.ApplicationInterface) error {
	router.Init(app.HttpHandler().(*gin.Engine))

	return nil
}

