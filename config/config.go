package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)


func LogFilePath() string {
	logFilePath := viper.GetString("logfile")
	if logFilePath[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		logFilePath = strings.Replace(logFilePath, "~", homeDir, 1)
	}
	return logFilePath
}

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
