package main

import (
	"rock/config"
	"os"
	"rock/cmd"
)

func main() {
	config.Init()
	os.OpenFile(config.LogFilePath(), os.O_CREATE, 0644)
	cmd.Execute()
}
