package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort        string
	AppSecret      string
	Debug          bool
	LogChannel     string
	ProfileEnabled bool

	// mysql
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	DBDatabase string

	// redis for session
	RedisDB       int
	RedisHost     string
	RedisPassword string
	RedisPort     string

	// minio
	MinioEndpoint     string
	MinioAccessKey    string
	MinioAccessSecret string
	MinioBucket       string
}

func Init(cfgFile string) *Config {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		viper.AddConfigPath(dir)
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	cfg := &Config{
		AppPort:           viper.GetString("app_port"),
		AppSecret:         viper.GetString("app_secret"),
		Debug:             viper.GetBool("debug"),
		LogChannel:        viper.GetString("log_channel"),
		ProfileEnabled:    viper.GetBool("profile_enabled"),
		DBHost:            viper.GetString("db_host"),
		DBPort:            viper.GetString("db_port"),
		DBUsername:        viper.GetString("db_username"),
		DBPassword:        viper.GetString("db_password"),
		DBDatabase:        viper.GetString("db_database"),
		RedisDB:           viper.GetInt("redis_db"),
		RedisHost:         viper.GetString("redis_host"),
		RedisPassword:     viper.GetString("redis_password"),
		RedisPort:         viper.GetString("redis_port"),
		MinioEndpoint:     viper.GetString("minio_endpoint"),
		MinioAccessKey:    viper.GetString("minio_access_key"),
		MinioAccessSecret: viper.GetString("minio_access_secret"),
		MinioBucket:       viper.GetString("minio_bucket"),
	}

	return cfg
}
