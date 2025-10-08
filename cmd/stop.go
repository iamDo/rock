package cmd

import (
	"fmt"
	"os"
	"rock/tracker"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop tracking a task",
	Args: cobra.NoArgs,
	Run: stopRun,
}

func stopRun(cmd *cobra.Command, args []string) {
	comment, err := cmd.Flags().GetString("comment")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = tracker.Stop(comment)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(stopCmd)
	stopCmd.Flags().StringP("comment", "c", "", "Add a comment to the start entry")
}
