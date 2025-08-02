package cmd

import (
	"fmt"
	"rock/tracker"
	"github.com/spf13/cobra"
	"os"
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

	err = tracker.Stop(comment, "rock.log")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(stopCmd)
	stopCmd.Flags().StringP("comment", "c", "", "Add a comment to the start entry")
}
