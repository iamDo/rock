package cmd

import (
	"fmt"
	"rock/tracker"
	"github.com/spf13/cobra"
	"os"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start tracking a task",
	Args: cobra.ExactArgs(1),
	Run: startRun,
}

func startRun(cmd *cobra.Command, args []string) {
	logFile := viper.GetString("logfile")
	comment, err := cmd.Flags().GetString("comment")
	ticket := args[0]
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tracker.Start(ticket, comment, logFile)
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("comment", "c", "", "Add a comment to the start entry")
}
