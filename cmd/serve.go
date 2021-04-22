package cmd

import (
	"github.com/DuC-cnZj/dota2app/pkg/app"
	"github.com/DuC-cnZj/dota2app/pkg/config"
	"github.com/DuC-cnZj/dota2app/pkg/dlog"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start dota2 server.",
	Run: func(cmd *cobra.Command, args []string) {
		app := app.NewApplication(config.Init(cfgFile))
		if err := app.Bootstrap(); err != nil {
			dlog.Fatal(err)
		}
		<-app.Run()
		app.Shutdown()
	},
}
