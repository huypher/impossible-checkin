package main

import (
	"fmt"
	ic "impossible-checkin"
	"impossible-checkin/process"
)

func main() {
	impossibleCheckIn := process.Process(ic.LogsFile)
	fmt.Println(impossibleCheckIn)
}
