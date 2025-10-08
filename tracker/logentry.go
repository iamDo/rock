package tracker

import (
	"fmt"
	"strings"
	"time"
	"unicode"
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
		comment = strings.TrimLeftFunc(splitLine[4], unicode.IsSpace)
	}

	now := time.Now()
	parsedTime, err := time.Parse("15:04", splitLine[0] + ":" + splitLine[1])
	if err != nil {
		return LogEntry{}, err
	}
	timeStamp := time.Date(now.Year(), now.Month(), now.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, now.Location())

	return LogEntry{
		Timestamp: timeStamp,
		Action: splitLine[2],
		Ticket: splitLine[3],
		Comment: comment,
	}, nil
}

func NewLogEntryNow(action string, ticket string, comment string) LogEntry {
	return LogEntry{
		Timestamp: time.Now(),
		Action: action,
		Ticket: ticket,
		Comment: comment,
	}
}
