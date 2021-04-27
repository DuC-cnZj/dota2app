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

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}
