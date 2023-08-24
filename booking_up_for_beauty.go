package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	time, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
	}
	return time
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	parsedDate, _ := time.Parse(layout, date)

	return time.Now().After(parsedDate)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	parsedDate, _ := time.Parse(layout, date)

	return parsedDate.Hour() >= 12 && parsedDate.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	time, _ := time.Parse(layout, date)
	return fmt.Sprintf("You have an appointment on %s, at %s.", time.Format("Monday, January 2, 2006"), time.Format("15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().UTC().Year()
	date := time.Date(currentYear, 9, 15, 0, 0, 0, 0, time.UTC)
	return date
}
