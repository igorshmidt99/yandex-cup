package main

import (
	"fmt"
	"time"
)

func main() {
	var alarms []Alarm
	for i := 0; i < 10; i++ {
		asStirng := string(i + 15000)
		duration, _ := time.ParseDuration(asStirng)
		alarm := Alarm{time.Now().Add(duration), duration}
		alarms = append(alarms, alarm)
	}

	fmt.Print(WakeUp(alarms, 5))
}
