// Package meetup implements method, which calculates exact day of month from vaguely entered date.
package meetup

import "time"

// WeekSchedule is designation for week order.
type WeekSchedule int

// First is order of the week.
const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

// Day returns day of month, which represents date, entered using arguments.
func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	firstWeekdayOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).Weekday()
	dateOfFirstWeekday := getDateOfWeekday(weekday, firstWeekdayOfMonth)
	if week < Teenth {
		return dateOfFirstWeekday + int(week)*7
	}
	var resultDay int
	if week == Teenth {
		for i := 1; ; i++ {
			resultDay = dateOfFirstWeekday + i*7
			if resultDay >= 13 {
				break
			}
		}
	} else if week == Last {
		dateOfLastMonthday := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
		for i := 1; ; i++ {
			resultDay = dateOfFirstWeekday + i*7
			if resultDay+7 > dateOfLastMonthday {
				break
			}
		}
	}
	return resultDay
}

func getDateOfWeekday(weekday time.Weekday, firstWeekdayOfMonth time.Weekday) int {
	result := int(weekday - firstWeekdayOfMonth + 1)
	if result <= 0 {
		result += 7
	}
	return result
}
