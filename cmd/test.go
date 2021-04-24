package cmd

import (
	"github.com/DuC-cnZj/dota2app/pkg/app"
	"github.com/DuC-cnZj/dota2app/pkg/config"
	"github.com/DuC-cnZj/dota2app/pkg/dlog"
	"github.com/DuC-cnZj/dota2app/pkg/models"
	"github.com/DuC-cnZj/dota2app/pkg/utils"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		app := app.NewApplication(config.Init(cfgFile))
		if err := app.Bootstrap(); err != nil {
			dlog.Fatal(err)
		}
		var u models.User

		utils.DB().First(&u)
		//hash, _ := utils.PasswordHash("12345")
		//u := &models.User{
		//	Name:     "duc",
		//	Email:    "1025434218@qq.com",
		//	Password: hash,
		//	Mobile:   "18888780080",
		//	Intro:    "hello everyone.",
		//}
		//utils.DB().Create(u)
		for _, file := range u.HistoryAvatars() {
			dlog.Info(file.GetFullPath(), file.FileableID)
		}
		app.Shutdown()
	},
}
