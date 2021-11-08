package goutils

import (
	"testing"
)

var (
	layout = "2006-01-02 15:04:05"
)

func TestIsTimeAfter(t *testing.T) {
	var tests = []struct {
		name string
		a, b string
		want bool
	}{
		{name: "diffrentDayFalse", a: "2021-10-02 15:04:05", b: "2006-01-02 15:04:05", want: false},
		{name: "diffrentDayTrue", a: "2006-01-02 15:04:05", b: "2021-10-02 15:04:05", want: true},
		{name: "diffrentSecondFalse", a: "2021-10-02 15:04:06", b: "2021-10-02 15:04:05", want: false},
		{name: "diffrentSecondTrue", a: "2021-10-02 15:04:04", b: "2021-10-02 15:04:05", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start, err := StringToTimeTime(layout, tt.a)
			if err != nil {
				t.Errorf("string to time.Time conversiton error: %v", err)
			}
			end, err := StringToTimeTime(layout, tt.b)
			if err != nil {
				t.Errorf("string to time.Time conversiton error: %v", err)
			}

			result := IsTimeAfter(start, end)
			if result != tt.want {
				t.Errorf("got %t, want %t", result, tt.want)
			}
		})
	}
}

func TestMilisecondsToTimeTime(t *testing.T) {
	var tests = []struct {
		name string
		got  int64
		want string
	}{
		{name: "milisecond to time.Time", got: 1633694979000, want: "2021-10-08 12:09:39"},
		{name: "milisecond to time.Time", got: 1633691379000, want: "2021-10-08 11:09:39"},
		{name: "milisecond to time.Time", got: 1633086579000, want: "2021-10-01 11:09:39"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MilisecondsToTimeTime(tt.got).Format(layout)
			if result != tt.want {
				t.Errorf("got %s, want %s", result, tt.want)
			}
		})
	}
}

func TestTimeTimeToString(t *testing.T) {
	var tests = []struct {
		name string
		got  string
		want string
	}{
		{name: "timeTimeToString 1", got: "2021-10-02 15:04:05", want: "2021-10-02 15:04:05 +0000 UTC"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toTimeTime, err := StringToTimeTime(layout, tt.got)
			if err != nil {
				t.Errorf("string to time.Time conversiton error: %v", err)
			}

			result := TimeTimeToString(toTimeTime)

			if result != tt.want {
				t.Errorf("got %s, want %s", result, tt.want)
			}
		})
	}
}
