package process

import (
	ic "impossible-checkin"
	"sort"
)

func SortLogsDataByLogID(logsData ic.LogsData) ic.LogsData {
	for _, logs := range logsData {
		sort.Slice(logs, func(i, j int) bool {
			return logs[i].LogID < logs[j].LogID
		})
	}

	return logsData
}
