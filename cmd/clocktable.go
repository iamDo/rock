package cmd

import (
	"fmt"
	"os"
	"rock/config"
	"rock/tracker"

	"github.com/spf13/cobra"
	"time"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type timeTicket struct {
	duration time.Duration
	ticket   string
}

var clocktableCmd = &cobra.Command{
	Use:   "clocktable",
	Short: "Displays today's clocktable",
	Run: clockTableRun,
}

func prettyDuration(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	if h == 0 {
		return fmt.Sprintf("%dm", m)
	}
	return fmt.Sprintf("%dh %dm", h, m)
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

	var rows [][]string
	for _, tt := range timeTickets {
		row := []string{tt.ticket, prettyDuration(tt.duration)}
		rows = append(rows, row)
	}

	var (
		purple    = lipgloss.Color("99")
		gray      = lipgloss.Color("245")
		lightGray = lipgloss.Color("241")

		headerStyle  = lipgloss.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)
		cellStyle    = lipgloss.NewStyle().Padding(0, 1).Width(14)
		oddRowStyle  = cellStyle.Foreground(gray)
		evenRowStyle = cellStyle.Foreground(lightGray)
	)

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(purple)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return headerStyle
			case row%2 == 0:
				return evenRowStyle
			default:
				return oddRowStyle
			}
		}).
		Headers("Ticket", "Duration").
		Rows(rows...)

	fmt.Println(t)
}

func init() {
	rootCmd.AddCommand(clocktableCmd)
}
