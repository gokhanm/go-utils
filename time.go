package goutils

import "time"

func TimeNowUTC() time.Time {
	return time.Now().UTC()
}

func IsTimeAfter(start, end time.Time) bool {
	return end.After(start)
}

func MilisecondsToTimeTime(ms int64) time.Time {
	return time.Unix(ms/1000, 0).UTC()
}
