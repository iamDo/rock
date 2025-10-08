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

	logData := strings.SplitSeq(string(dat[:]), "\n")
	for log := range logData {
		l, err := ParseLogEntry(log)
		if err == nil {
			fmt.Println(l.Timestamp)
		}
	}

	return "", nil
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
