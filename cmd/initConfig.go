package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var initConfigCmd = &cobra.Command{
	Use:   "initConfig",
	Short: "初始化配置文件",
	Run: func(cmd *cobra.Command, args []string) {
		getwd, _ := os.Getwd()
		configPath := getwd + "/config.yaml"
		if len(configYaml) > 0 {
			if Exists(configPath) {
				log.Printf("文件 %s 已存在", configPath)
				return
			}
			f, err := os.Create(configPath)
			if err != nil {
				log.Println(err)
				return
			}
			defer f.Close()
			f.Write(configYaml)

			log.Printf("文件 %s 创建成功", configPath)
		}
	},
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}