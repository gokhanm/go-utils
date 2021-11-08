package goutils

import (
	"time"
)

func TimeNowUTC() time.Time {
	return time.Now().UTC()
}

// IsTimeAfter Is the b comes after the a?
func IsTimeAfter(a, b time.Time) bool {
	return b.After(a)
}

// MilisecondsToTimeTime converts int64 to time.Time. location set to UTC
func MilisecondsToTimeTime(ms int64) time.Time {
	return time.Unix(ms/1000, 0).UTC()
}

func StringToTimeTime(format, day string) (time.Time, error) {
	t, err := time.Parse(format, day)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func TimeTimeToString(t time.Time) string {
	return t.String()
}
