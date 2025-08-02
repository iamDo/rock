package cmd

import (
	"fmt"
	"rock/tracker"
	"rock/config"
	"github.com/spf13/cobra"
	"os"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start tracking a task",
	Args: cobra.ExactArgs(1),
	Run: startRun,
}

func startRun(cmd *cobra.Command, args []string) {
	logFile := config.LogFilePath()
	comment, err := cmd.Flags().GetString("comment")
	ticket := args[0]
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = tracker.Start(ticket, comment, logFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("comment", "c", "", "Add a comment to the start entry")
}
