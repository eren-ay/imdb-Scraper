package models

import "time"

type Date struct {
	Year  int
	Month int
	Day   int
}

func (date *Date) Age() int {
	currentTime := time.Now()

	year := currentTime.Year()
	personAge := year - date.Year
	return personAge
}
