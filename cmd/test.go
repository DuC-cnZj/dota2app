package cmd

import (
	"context"
	"fmt"
	"github.com/DuC-cnZj/dota2app/pkg/app"
	"github.com/DuC-cnZj/dota2app/pkg/config"
	"github.com/DuC-cnZj/dota2app/pkg/contracts"
	"github.com/DuC-cnZj/dota2app/pkg/dlog"
	"github.com/DuC-cnZj/dota2app/pkg/models"
	"github.com/DuC-cnZj/dota2app/pkg/utils"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
	"sync"
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
		var files []*models.File
		utils.DB().Unscoped().Find(&files)
		wg:=sync.WaitGroup{}
		wg.Add(len(files))
		for _, file := range files {
			go func(file *models.File) {
				defer wg.Done()
				obj, err := file.ToMinioUploadInfo()
				if err != nil {
					dlog.Error(err)
					return
				}
				fmt.Println(err,obj)
				dlog.Info(obj)
				if err := app.FileManager().(contracts.WithMinio).MinioClient().RemoveObject(context.Background(), app.Config().MinioBucket, obj.Key, minio.RemoveObjectOptions{}); err != nil {
					dlog.Error(err)
				}
			}(file)
		}
		wg.Wait()
		//for _, file := range u.HistoryAvatars() {
		//	utils.DB().Delete(&file)
		//	//dlog.Info(file.GetFullPath(), file.FileableID)
		//}
		//for _, file := range u.HistoryBackgrounds() {
			//utils.DB().Delete(&file)
			// obj, _ := file.ToMinioUploadInfo()
			//dlog.Info(file.GetFullPath(), file.FileableID)
		//}
		app.Shutdown()
	},
}
