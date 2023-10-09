package main

import (
	"fmt"
	"time"
)

func changeTimeZone(dst bool, zone string) *time.Location {
	//time difference due to daylight savings
	tDiff := 0
	var loc *time.Location
	if dst {
		tDiff = 1
	}

	switch zone {
	case "EST":
		loc = time.FixedZone("EST", (-5+tDiff)*3600)
	case "PST":
		loc = time.FixedZone("PST", (-8+tDiff)*3600)
	case "CST":
		loc = time.FixedZone("CST", (-6+tDiff)*3600)
	default:
		loc = time.FixedZone("UTC", 0)
	}
	return loc
}

func main() {
	//time format
	var format string
	fmt.Println("What format would you'd like the time")
	fmt.Println("EST, PST, OR CST: ")
	fmt.Scan(&format)

	currTime := time.Now().UTC()
	daylight := time.Now().IsDST()

	//The Newly Formatted Time
	newLoc := changeTimeZone(daylight, format)
	newTime := currTime.In(newLoc).Format(time.RFC1123)
	fmt.Println(newTime)
}
