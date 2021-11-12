package goutils

import (
	"reflect"
	"testing"
	"time"
)

func TestIsTimeAfter(t *testing.T) {
	type args struct {
		a time.Time
		b time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "false", args: args{a: time.Date(9999, 4, 12, 23, 20, 50, 520*1e6, time.UTC), b: time.Date(1996, 12, 19, 16, 39, 57, 0, time.UTC)}, want: false},
		{name: "true", args: args{a: time.Date(1996, 12, 19, 16, 39, 57, 0, time.UTC), b: time.Date(9999, 4, 12, 23, 20, 50, 520*1e6, time.UTC)}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTimeAfter(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IsTimeAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMilisecondsToTimeTime(t *testing.T) {
	type args struct {
		ms int64
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{name: "1221681866000", args: args{1221681866000}, want: time.Date(2008, 9, 17, 20, 4, 26, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MilisecondsToTimeTime(tt.args.ms); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MilisecondsToTimeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToTimeTime(t *testing.T) {
	type args struct {
		format string
		day    string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{name: "RFC3339 format", args: args{format: time.RFC3339, day: "2008-09-17T20:04:26Z"}, want: time.Date(2008, 9, 17, 20, 4, 26, 0, time.UTC), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToTimeTime(tt.args.format, tt.args.day)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToTimeTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToTimeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeTimeToString(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "toString1", args: args{time.Date(9999, 4, 12, 23, 20, 50, 520*1e6, time.UTC)}, want: "9999-04-12 23:20:50.52 +0000 UTC"},
		{name: "toString2", args: args{time.Date(1996, 12, 19, 16, 39, 57, 0, time.UTC)}, want: "1996-12-19 16:39:57 +0000 UTC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeTimeToString(tt.args.t); got != tt.want {
				t.Errorf("TimeTimeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
