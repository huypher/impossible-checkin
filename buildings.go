package impossible_checkin

import (
	"time"
)

// I don't have enough time for this test so I decide hardcode buildings data here

type Building int

const (
	BuildingA Building = iota
	BuildingB
	BuildingC
)

func (b Building) String() string {
	return [...]string{"BuildingA", "BuildingB", "BuildingC"}[b]
}

var BuildingMap map[string]Building

var BuildingTravel [][]time.Duration

func init() {
	BuildingMap = map[string]Building{
		BuildingA.String(): BuildingA,
		BuildingB.String(): BuildingB,
		BuildingC.String(): BuildingC,
	}

	BuildingTravel = make([][]time.Duration, len(BuildingMap))
	for i := 0; i < len(BuildingMap); i++ {
		BuildingTravel[i] = make([]time.Duration, len(BuildingMap))
	}

	BuildingTravel[BuildingA][BuildingB] = 10 * time.Minute
	BuildingTravel[BuildingA][BuildingC] = 20 * time.Minute
	BuildingTravel[BuildingB][BuildingC] = 50 * time.Minute
}

func GetTravelTime(from string, to string) (bool, time.Duration) {
	var fromBuilding Building
	var toBuilding Building
	if id, ok := BuildingMap[from]; !ok {
		return false, 0
	} else {
		fromBuilding = id
	}

	if id, ok := BuildingMap[to]; !ok {
		return false, 0
	} else {
		toBuilding = id
	}

	travelTime := BuildingTravel[fromBuilding][toBuilding]
	if travelTime == 0 {
		return false, 0
	}

	return true, travelTime
}
