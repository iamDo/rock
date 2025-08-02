package config

import (
	"os"
	"github.com/spf13/viper"
	"path/filepath"
)


func Init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("couldn't get config dir: " + err.Error())
	}

	viper.SetConfigName("rock")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Join(homeDir, ".config", "rock"))
	viper.AddConfigPath(".")
	viper.SetDefault("logfile", "./rock.log")

	err = viper.ReadInConfig()
	if err != nil {
		panic("could not read config" + err.Error())
	}
}
