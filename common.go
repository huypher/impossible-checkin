package impossible_checkin

import "time"

const (
	Sep        = " "
	LogDataLen = 5
	TimeLayout = "2006-01-02T15:04:05"

	LogsFile = "./logs.txt"
)

type ImpossibleCheckIn []string
type LogsData map[int][]Log
type InvalidLogsData []string

type Log struct {
	LogID      string
	EmployeeID int
	Location   string
	CheckIn    time.Time
	CheckOut   time.Time
}
