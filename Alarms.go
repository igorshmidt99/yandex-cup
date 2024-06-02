package main

import (
	"slices"
	"time"
)

type Alarm struct {
	timestamp   time.Time
	periodicity time.Duration
}

// Текущее время - это время последнего будильника. Оно возвращается

// запомнил timestamp последнего
// как только достал первый будильник, убавляю счетчик alarmCount. Если alarmCount == 0, то возвращаю timestamp
// затем timestamp + periodicity последнего.
// отсортировал
// Если последний запомненный timestamp == timestamp первого элемента во вновь отсортированном в списке,
// то тогда я пропускаю этот будильник и беру следующий (конфликт)

func WakeUp(alarms []Alarm, alarmCount int) time.Time {
	// Самый первый шаг
	slices.SortStableFunc(alarms, compare)
	headAlarm := getFirst(alarms)
	alarmCount--
	lastAlarmTime := headAlarm.timestamp
	headAlarm.timestamp = lastAlarmTime.Add(headAlarm.periodicity)
	return onTime(alarms, alarmCount, lastAlarmTime)
}

func onTime(alarms []Alarm, alarmCount int, lastAlarmTime time.Time) time.Time {
	if alarmCount == 0 {
		return lastAlarmTime
	}

	slices.SortStableFunc(alarms, compare)
	head := getFirst(alarms)
	if head.timestamp == lastAlarmTime {
		head = getWithoutConflict(alarms, lastAlarmTime)
	}
	lastAlarmTime = head.timestamp
	head.timestamp = lastAlarmTime.Add(head.periodicity)
	alarmCount--
	return onTime(alarms, alarmCount, lastAlarmTime)
}

func getWithoutConflict(alarms []Alarm, lastAlarmTime time.Time) Alarm {
	for _, alarm := range alarms {
		if alarm.timestamp != lastAlarmTime {
			return alarm
		}
	}
	return alarms[0]
}

func getFirst(alarms []Alarm) Alarm {
	return alarms[0]
}

func compare(a, b Alarm) int {
	return a.timestamp.Compare(b.timestamp)
}
