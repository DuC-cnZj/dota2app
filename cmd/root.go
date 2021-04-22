package cmd

import (
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "dota2 app.",
}

func Execute() {
	rootCmd.AddCommand(serveCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $DIR/config.yaml)")
	rootCmd.PersistentFlags().BoolP("debug", "", true, "debug mode.")
	rootCmd.PersistentFlags().StringP("app_port", "", "5000", "app port.")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("app_port", rootCmd.PersistentFlags().Lookup("app_port"))
}

