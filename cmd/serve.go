package cmd

import (
	"rock/server"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the rock http server",
	Run: serveRun,
}

func serveRun(cmd *cobra.Command, args []string) {
	server.Serve()
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
