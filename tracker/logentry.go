package tracker

import (
	"fmt"
	"strings"
	"time"
)

const logEntryFormat = "%s:%s:%s: "
const timeStampFormat = "15:04"

type LogEntry struct {
	Timestamp time.Time
	Action    string
	Ticket	  string
	Comment   string
}

type LogEntryParseError struct {
	message string
}

func (e LogEntryParseError) Error() string {
	return e.message
}

func (l LogEntry) String() string {
	return fmt.Sprintf(logEntryFormat, l.Timestamp.Format(timeStampFormat), l.Action, l.Ticket) + l.Comment
}

func ParseLogEntry(l string) (LogEntry, error) {
	splitLine := strings.SplitN(l, ":", 5)

	if len(splitLine) != 4 && len(splitLine) != 5 {
		return LogEntry{}, LogEntryParseError{fmt.Sprintf("line %s cannot be split", l)}
	}

	comment := ""
	if (len(splitLine) > 4) {
		comment = splitLine[4]
	}

	timeStamp, err := time.Parse("15:04", splitLine[0] + ":" + splitLine[1])
	if err != nil {
		return LogEntry{}, err
	}

	return LogEntry{
		Timestamp: timeStamp,
		Action: splitLine[2],
		Ticket: splitLine[3],
		Comment: comment,
	}, nil
}
