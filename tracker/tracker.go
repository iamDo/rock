package tracker

import (
	"fmt"
	"os"
	"strings"
)

func Start(ticket string, comment string, logFile string) error {
	logEntry := NewLogEntryNow("start", ticket, comment)
	lastEntry, err := lastLogEntry(logFile)
	if err != nil {
		return err
	}

	if lastEntry.Action == "start" {
		stopLogEntry := NewLogEntryNow("stop", lastEntry.Ticket, "automatically clocked out")
		err := writeLog(stopLogEntry, logFile)
		if err != nil {
			return err
		}
	}

	return writeLog(logEntry, logFile)
}

func Stop(comment string, logFile string) error {
	lastEntry, err := lastLogEntry(logFile)
	if err != nil {
		return err
	}

	if lastEntry.Action == "stop" {
		return fmt.Errorf("Nothing to clock out from")
	}

	ticket := lastEntry.Ticket

	logEntry := NewLogEntryNow("stop", ticket, comment)
	return writeLog(logEntry, logFile)
}

func lastLogEntry(logFile string) (LogEntry, error) {
	logEntries, err := getLogEntries(logFile)
	if err != nil {
		return LogEntry{}, err
	}

	if logEntries == nil {
		return LogEntry{}, nil
	}

	return logEntries[len(logEntries) - 1], nil
}

func getLogEntries(logFile string) ([]LogEntry, error) {
	dat, err := os.ReadFile(logFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	if len(dat) == 0 {
		return nil, nil
	}

	logData := strings.SplitSeq(string(dat[:]), "\n")
	var logEntries []LogEntry
	for log := range logData {
		l, err := ParseLogEntry(log)
		if err == nil {
			logEntries = append(logEntries, l)
		}
	}
	return logEntries, nil
}

func GetClockedState(ticket string, logFile string) (string, error) {
	logEntries, err := getLogEntries(logFile)
	if err != nil {
		return "", err
	}
	var matchingEntries []LogEntry
	for _, log := range logEntries {
		if log.Ticket == ticket {
			matchingEntries = append(matchingEntries, log)
		}
	}
	lastEntry := matchingEntries[len(matchingEntries) - 1]

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
