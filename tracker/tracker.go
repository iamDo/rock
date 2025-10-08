package tracker

import (
	"fmt"
	"os"
	"strings"
)

func Start(ticket string, comment string, logFile string) error {
	logEntry := NewLogEntryNow("start", ticket, comment)
	return writeLog(logEntry, logFile)
}

func Stop(ticket string, comment string, logFile string) error {
	logEntry := NewLogEntryNow("start", ticket, comment)
	return writeLog(logEntry, logFile)
}

func GetClockedState(ticket string, logFile string) (string, error) {
	dat, err := os.ReadFile(logFile)
	if err != nil {
		return "", fmt.Errorf("failed to open log file: %w", err)
	}

	logData := strings.Split(string(dat[:]), "\n")
	var logEntries []LogEntry
	for _, log := range logData {
		l, err := ParseLogEntry(log)
		if err == nil && l.Ticket == ticket {
			logEntries = append(logEntries, l)
		}
	}
	lastEntry := logEntries[len(logEntries) - 1]

	return lastEntry.Action, nil
}

func writeLog(logEntry LogEntry, logFile string) error {
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer f.Close()

	_, err = f.WriteString(logEntry.String() + "\n")
	return err
}
