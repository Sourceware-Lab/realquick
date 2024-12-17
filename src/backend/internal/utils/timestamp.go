package utils

import "time"

type TimeStamp struct {
	Start time.Time
	End   time.Time
}

func IsSameDate(start, end time.Time) bool {
	return false // todo, fix this
}

func IsDateAfter(start, end time.Time) bool {
	return false // todo, fix this
}
