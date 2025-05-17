package main

import "time"

func main() {

	//check day of week
	now := time.Now()
	dayOfWeek := now.Weekday()

	//check if day of week is saturday or sunday
	if dayOfWeek == time.Saturday || dayOfWeek == time.Sunday {
		// skip this day
	}

	// check clock
	hour := now.Hour()

	// fixme : learn about market opening and closing times
	//if not between 12am (00:00) and 12pm (12:00) skip this day
	if hour >= 12 {
		// skip this day
	}

}
