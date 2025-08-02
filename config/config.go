package config

import (
	"os"
	"github.com/spf13/viper"
	"path/filepath"
)


func Init() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic("couldn't get config dir: " + err.Error())
	}

	viper.SetConfigName("rock")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Join(configDir, "rock"))
	viper.AddConfigPath(".")
	viper.SetDefault("logfile", "./rock.log")

	err = viper.ReadInConfig()
	if err != nil {
		panic("could not read config" + err.Error())
	}
}
