package tracker

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const logFormat = "%s:%s:%s: "

func Start(ticket string, comment string, logFile string) error {
	logTime := time.Now().Format("15:04")
	logLine := fmt.Sprintf(logFormat, logTime, "start", ticket) + comment + "\n"
	return writeLog(logLine, logFile)
}

func Stop(ticket string, comment string, logFile string) error {
	logTime := time.Now().Format("15:04")
	logLine := fmt.Sprintf(logFormat, logTime, "stop", ticket) + comment + "\n"
	return writeLog(logLine, logFile)
}

func GetClockedState(ticket string, logFile string) (string, error) {
	dat, err := os.ReadFile(logFile)
	if err != nil {
		return "", fmt.Errorf("failed to open log file: %w", err)
	}

	logData := strings.SplitSeq(string(dat[:]), "\n")
	for log := range logData {
		l, err := ParseLogEntry(log)
		if err != nil {
			return "", fmt.Errorf("failed to parse log line: %w", err)
		}

		fmt.Println(l.Timestamp)
	}

	return "", nil
}

func writeLog(logLine string, logFile string) error {
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer f.Close()

	_, err = f.WriteString(logLine)
	return err
}
