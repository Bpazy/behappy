package btime

import "time"

func GetLastWeek() (start time.Time, end time.Time) {
	monday := GetThisWeekMonday()
	return monday.AddDate(0, 0, -7), monday
}

func GetThisWeekMonday() time.Time {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	return RoundDate(now).AddDate(0, 0, offset)
}

// RoundDate Remove time.Time's hour, min, sec
func RoundDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
