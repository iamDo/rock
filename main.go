package main

import (
	"rock/config"
	"os"
	"rock/cmd"
)

func main() {
	config.Init()
	f, err := os.OpenFile(config.LogFilePath(), os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	f.Close()
	cmd.Execute()
}
