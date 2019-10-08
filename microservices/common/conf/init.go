package conf

import (
	"strings"

	"github.com/spf13/viper"
)

type WebConfig struct {
	ListenAddr string
}

type MongoConfig struct {
	Uri string
}

type DbConfig struct {
	Mongo MongoConfig
}

type Config struct {
	Web WebConfig
	Db  DbConfig
}

func initViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		//panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
}

func Init() {
	initViper() //using conf file and env
}

func IsDebugEnv() bool {
	return viper.GetString("mode") == "debug"
}
