package booking

import (
	"fmt"
	"time"
)

/*
	Year	2006 ; 06
	Month	Jan ; January ; 01 ; 1
	Day	02 ; 2 ; _2 (For preceding 0)
	Weekday	Mon ; Monday
	Hour	15 ( 24 hour time format ) ; 3 ; 03 (AM or PM)
	Minute	04 ; 4
	Second	05 ; 5
	AM/PM Mark	PM
	Day of Year	002 ; __2
*/

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// MM/dd/yyyy hh:MM:ss
	layout := "1/2/2006 15:04:05"
	dt, err := time.Parse(layout, date)

	if err != nil {
		panic(err)
	}
	return dt
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	dt, err := time.Parse(layout, date)

	if err != nil {
		panic(err)
	}
	return dt.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	dt, err := time.Parse(layout, date)

	if err != nil {
		panic(err)
	}

	hour := dt.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layoutParser := "1/2/2006 15:04:05"
	layoutConvert := "Monday, January 2, 2006, at 15:04"

	dt, err := time.Parse(layoutParser, date)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("You have an appointment on %s.", dt.Format(layoutConvert))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	date := fmt.Sprintf("%d-09-15 00:00:00", time.Now().Year())
	layout := "2006-01-02 15:04:05"

	dt, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	return dt
}
