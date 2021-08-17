package process

import (
	ic "impossible-checkin"
	"impossible-checkin/read"
)

func Process(logsFile string) ic.ImpossibleCheckIn {
	logsData, _, _ := read.ReadLogsData(logsFile)
	logsData = SortLogsDataByLogID(logsData)

	impossibleCheckIn := make(ic.ImpossibleCheckIn, 0)

	for _, logs := range logsData {
		var previousValid *ic.Log
		for idx, l := range logs {
			if previousValid == nil {
				previousValid = &logs[idx]
				continue
			}

			if previousValid.CheckOut.IsZero() {
				impossibleCheckIn = append(impossibleCheckIn, l.LogID)
				continue
			}

			if l.CheckIn.Before(previousValid.CheckOut) {
				impossibleCheckIn = append(impossibleCheckIn, l.LogID)
				continue
			}

			if l.Location == previousValid.Location {
				previousValid = &logs[idx]
				continue
			}

			isTravel, travelTime := ic.GetTravelTime(previousValid.Location, l.Location)
			if !isTravel {
				impossibleCheckIn = append(impossibleCheckIn, l.LogID)
				continue
			}

			if (l.CheckIn.Add(-travelTime)).Before(previousValid.CheckOut) {
				impossibleCheckIn = append(impossibleCheckIn, l.LogID)
				continue
			}

			previousValid = &logs[idx]
		}
	}

	return impossibleCheckIn
}
