package main

import (
	"rock/config"
	"rock/cmd"
)

func main() {
	config.Init()
	cmd.Execute()
}
