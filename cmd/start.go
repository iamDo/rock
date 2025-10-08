package cmd

import (
	"fmt"
	"strings"

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
	comment, err := cmd.Flags().GetString("comment")
	ticket := args[0]
	var out strings.Builder
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	out.WriteString("Clocked IN on ")
	out.WriteString(ticket)
	if comment != "" {
		out.WriteString(": " + comment)
	}

	fmt.Println(out.String())

}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("comment", "c", "", "Add a comment to the start entry")
}
