package random

import (
	"time"
	"math/rand"
	"fmt"
)

const dateTimeFormat = "2006-01-02T15:04:05Z"
const dateFormat = "2006-01-02"

func DateTimeBefore(orig time.Time, minDays int, maxDays int) time.Time {
	var h, m, s int
	rDays := Int(minDays, maxDays)
	rTime := orig.AddDate(0, 0, -rDays)

	if rDays == 0 {
		h = rand.Intn(orig.Hour())
		m = rand.Intn(orig.Minute())
		s = rand.Intn(orig.Second())
	} else {
		h = rand.Intn(23)
		m = rand.Intn(59)
		s = rand.Intn(59)
	}

	rTime, _ = time.Parse(dateTimeFormat, fmt.Sprintf("%sT%d:%d:%dZ", rTime.Format(dateFormat), h, m, s))

	return rTime
}

func TimeBefore(orig time.Time) time.Time {
	rand.Seed(time.Now().UnixNano())
	h := rand.Intn(orig.Hour())
	m := rand.Intn(orig.Minute())
	s := rand.Intn(orig.Second())

	rTime, _ := time.Parse(dateTimeFormat, fmt.Sprintf("%sT%d:%d:%dZ", orig.Format(dateFormat), h, m, s))

	return rTime
}

func Int(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max - min)
}
