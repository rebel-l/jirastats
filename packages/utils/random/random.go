package random

import (
	"fmt"
	"math/rand"
	"time"
)

const dateTimeFormat = "2006-01-02T15:04:05Z"
const dateFormat = "2006-01-02"

var seedSet bool = false

func DateTimeBefore(orig time.Time, minDays int, maxDays int) time.Time {
	var h, m, s int
	rDays := Int(minDays, maxDays)
	rTime := orig.AddDate(0, 0, -rDays)

	if rDays == 0 {
		h, m, s = getTime(orig.Hour(), orig.Minute(), orig.Second())
	} else {
		h, m, s = getTime(23, 59, 59)
	}

	rTime, _ = time.Parse(dateTimeFormat, fmt.Sprintf("%sT%02d:%02d:%02dZ", rTime.Format(dateFormat), h, m, s))

	return rTime
}

func TimeBefore(orig time.Time) time.Time {
	if seedSet == false {
		rand.Seed(time.Now().UnixNano())
		seedSet = true
	}
	
	h, m, s := getTime(orig.Hour(), orig.Minute(), orig.Second())

	rTime, _ := time.Parse(dateTimeFormat, fmt.Sprintf("%sT%02d:%02d:%02dZ", orig.Format(dateFormat), h, m, s))

	return rTime
}

func getTime(maxHour int, maxMin int, maxSec int) (hour int, min int, sec int) {
	if maxHour > 0 && maxHour < 24 {
		hour = rand.Intn(maxHour)
	}

	if maxMin > 0 && maxMin < 60 {
		min = rand.Intn(maxMin)
	}

	if maxSec > 0 && maxSec < 60 {
		sec = rand.Intn(maxSec)
	}
	return
}

func Int(min int, max int) int {
	if seedSet == false {
		rand.Seed(time.Now().UnixNano())
		seedSet = true
	}

	return min + rand.Intn(max-min)
}
