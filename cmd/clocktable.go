package cmd

import (
	"fmt"
	"os"
	"rock/config"
	"rock/tracker"

	"github.com/spf13/cobra"
	"time"
)

type timeTicket struct {
	duration time.Time
	ticket   string
}

var clocktableCmd = &cobra.Command{
	Use:   "clocktable",
	Short: "Displays today's clocktable",
	Run: clockTableRun,
}

func clockTableRun(cmd *cobra.Command, args []string) {
	uniqueTickets, err := tracker.UniqueTickets(config.LogFilePath())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var timeTickets []timeTicket
	for _, t := range uniqueTickets {
		ts, err := tracker.TimeSpent(t, config.LogFilePath())
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		timeTickets = append(timeTickets, timeTicket{ts, t})
	}

	for _, tt := range timeTickets {
		fmt.Println(fmt.Sprintf("[%s] %s", tt.ticket, tt.duration.Format("15:04")))
	}
}

func init() {
	rootCmd.AddCommand(clocktableCmd)
}
