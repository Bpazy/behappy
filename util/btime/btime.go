package btime

import "time"

func GetLastWeekRange() (start time.Time, end time.Time) {
	monday := GetCurrentWeekMonday()
	return monday.AddDate(0, 0, -7), monday
}

func GetCurrentWeekRange() (start time.Time, end time.Time) {
	monday := GetCurrentWeekMonday()
	return monday, monday.AddDate(0, 0, 7)
}

func GetCurrentWeekMonday() time.Time {
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
