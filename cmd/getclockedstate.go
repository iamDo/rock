package cmd

import (
	"fmt"
	"os"
	"rock/tracker"

	"github.com/spf13/cobra"
)

var getclockedstateCmd = &cobra.Command{
	Use:   "getclockedstate",
	Short: "Get the clocked state of a given ticket",
	Args: cobra.ExactArgs(1),
	Run: getClockedStateRun,
}

func getClockedStateRun(cmd *cobra.Command, args []string) {
	state, err := tracker.GetClockedState(args[0], "rock.log")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%s", state)
}

func init() {
	rootCmd.AddCommand(getclockedstateCmd)
}
