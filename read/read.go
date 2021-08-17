package read

import (
	"bufio"
	ic "impossible-checkin"
	"os"
	"strconv"
	"strings"
	"time"
)

// We assume that if employee was not yet checked out from another session,
// check-out time in that session is 0001-01-01T00:00:00 in logs file
func ReadLogsData(logsFile string) (ic.LogsData, ic.InvalidLogsData, error) {
	file, err := os.Open(logsFile)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	logsData := make(ic.LogsData)
	invalidLogsData := make(ic.InvalidLogsData, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitedLine := strings.Split(line, ic.Sep)
		if len(splitedLine) != ic.LogDataLen {
			invalidLogsData = append(invalidLogsData, line)
			continue
		}

		var employeeID int
		var checkIn time.Time
		var checkOut time.Time

		if id, err := strconv.Atoi(strings.TrimSpace(splitedLine[1])); err != nil {
			invalidLogsData = append(invalidLogsData, line)
			continue
		} else {
			employeeID = id
		}

		if checkInTime, err := time.Parse(ic.TimeLayout, strings.TrimSpace(splitedLine[3])); err != nil {
			invalidLogsData = append(invalidLogsData, line)
			continue
		} else {
			checkIn = checkInTime
		}

		if checkOutTime, err := time.Parse(ic.TimeLayout, strings.TrimSpace(splitedLine[4])); err != nil {
			invalidLogsData = append(invalidLogsData, line)
			continue
		} else {
			checkOut = checkOutTime
		}

		log := ic.Log{
			LogID:      strings.TrimSpace(splitedLine[0]),
			EmployeeID: employeeID,
			Location:   strings.TrimSpace(splitedLine[2]),
			CheckIn:    checkIn,
			CheckOut:   checkOut,
		}

		if logs, ok := logsData[employeeID]; ok {
			logsData[employeeID] = append(logs, log)
			continue
		}

		logsData[employeeID] = []ic.Log{
			{
				LogID:      splitedLine[0],
				EmployeeID: employeeID,
				Location:   splitedLine[2],
				CheckIn:    checkIn,
				CheckOut:   checkOut,
			},
		}
	}

	return logsData, invalidLogsData, nil
}
