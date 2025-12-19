package booking

import (
    "time"
    "fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
	t, _ := time.Parse(layout, date)

	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	now := time.Now()
	layout := "January 2, 2006 15:04:05"
	dt, _ := time.Parse(layout, date)
	return dt.Before(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	dt, _ := time.Parse(layout, date)
	minT := time.Date(dt.Year(), dt.Month(), dt.Day(), 12, 0, 0, 0, dt.Location())
	maxT := time.Date(dt.Year(), dt.Month(), dt.Day(), 18, 0, 0, 0, dt.Location())
	return dt.After(minT) && dt.Before(maxT)
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	dt, _ := time.Parse(layout, date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", dt.Weekday().String(), dt.Month().String(), dt.Day(), dt.Year(), dt.Hour(), dt.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	now := time.Now()
	ad := time.Date(now.Year(), 9, 15, 0, 0, 0, 0, time.UTC)
	return ad
}
