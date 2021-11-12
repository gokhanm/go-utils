package goutils

import (
	"database/sql"
	"reflect"
	"testing"
	"time"
)

func TestStringToInt64(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{name: "no error", args: args{"10"}, want: 10, wantErr: false},
		{name: "want error", args: args{""}, want: 0, wantErr: true},
		{name: "-1", args: args{"-1"}, want: -1, wantErr: false},
		{name: "-12345", args: args{"-12345"}, want: -12345, wantErr: false},
		{name: "bigInt", args: args{"9223372036854775807"}, want: 1<<63 - 1, wantErr: false},
		{name: "error range", args: args{"-9223372036854775809"}, want: -1 << 63, wantErr: true},
		{name: "error syntax", args: args{"-1_2_3_4_5"}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToInt64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToBoolen(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "no error", args: args{"true"}, want: true, wantErr: false},
		{name: "err syntax", args: args{""}, want: false, wantErr: true},
		{name: "err syntax", args: args{"asd"}, want: false, wantErr: true},
		{name: "0 no error", args: args{"0"}, want: false, wantErr: false},
		{name: "FALSE", args: args{"FALSE"}, want: false, wantErr: false},
		{name: "False", args: args{"False"}, want: false, wantErr: false},
		{name: "TRUE", args: args{"TRUE"}, want: true, wantErr: false},
		{name: "True", args: args{"True"}, want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToBoolen(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToBoolen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringToBoolen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToInt(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "1", args: args{"1"}, want: 1, wantErr: false},
		{name: "-1", args: args{"-1"}, want: -1, wantErr: false},
		{name: "-0", args: args{"-0"}, want: 0, wantErr: false},
		{name: "errs syntax", args: args{""}, want: 0, wantErr: true},
		{name: "errs range", args: args{"9223372036854775809"}, want: 1<<63 - 1, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToInt(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToFloat64(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "0.1", args: args{"0.1"}, want: 0.1, wantErr: false},
		{name: "-1", args: args{"-1"}, want: -1.0, wantErr: false},
		{name: "-1", args: args{"-1"}, want: -1.0, wantErr: false},
		{name: "-0", args: args{"-0"}, want: -0, wantErr: false},
		{name: "err syntax", args: args{""}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToFloat64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64ToString(t *testing.T) {
	type args struct {
		value int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "9223372036854775807", args: args{1<<63 - 1}, want: "9223372036854775807"},
		{name: "10000", args: args{10000}, want: "10000"},
		{name: "1", args: args{1}, want: "1"},
		{name: "-1", args: args{-1}, want: "-1"},
		{name: "0", args: args{-0}, want: "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64ToString(tt.args.value); got != tt.want {
				t.Errorf("Int64ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64ToString(t *testing.T) {
	type args struct {
		key float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0.9", args: args{0.9}, want: "0.9"},
		{name: "0.5", args: args{0.5}, want: "0.5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64ToString(tt.args.key); got != tt.want {
				t.Errorf("Float64ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqlNullString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{name: "1", args: args{"1"}, want: sql.NullString{String: "1", Valid: true}},
		{name: "empty", args: args{""}, want: sql.NullString{String: "", Valid: false}},
		{name: "0", args: args{"0"}, want: sql.NullString{String: "0", Valid: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SqlNullString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SqlNullString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeToSqlNullTime(t *testing.T) {
	t0 := time.Time{}
	t1 := time.Date(2000, 1, 1, 8, 9, 10, 11, time.UTC)

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want sql.NullTime
	}{
		{name: "t0", args: args{t0}, want: sql.NullTime{Time: t0, Valid: false}},
		{name: "t1", args: args{t1}, want: sql.NullTime{Time: t1, Valid: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeToSqlNullTime(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeToSqlNullTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64ToSqlNullInt64(t *testing.T) {
	type args struct {
		t int64
	}
	tests := []struct {
		name string
		args args
		want sql.NullInt64
	}{
		{name: "30", args: args{30}, want: sql.NullInt64{Int64: 30, Valid: true}},
		{name: "-22", args: args{-22}, want: sql.NullInt64{Int64: -22, Valid: true}},
		{name: "empty", args: args{}, want: sql.NullInt64{Int64: 0, Valid: false}},
		{name: "zero", args: args{0}, want: sql.NullInt64{Int64: 0, Valid: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64ToSqlNullInt64(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64ToSqlNullInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64ToSqlNullFloat64(t *testing.T) {
	type args struct {
		t float64
	}
	tests := []struct {
		name string
		args args
		want sql.NullFloat64
	}{
		{name: "-22.2", args: args{-22.2}, want: sql.NullFloat64{Float64: -22.2, Valid: true}},
		{name: "empty", args: args{}, want: sql.NullFloat64{Float64: 0, Valid: false}},
		{name: "zero", args: args{0}, want: sql.NullFloat64{Float64: 0, Valid: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64ToSqlNullFloat64(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64ToSqlNullFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolenToSqlNullBoolen(t *testing.T) {
	type args struct {
		t bool
	}
	tests := []struct {
		name string
		args args
		want sql.NullBool
	}{
		{name: "false", args: args{false}, want: sql.NullBool{Bool: false, Valid: false}},
		{name: "true", args: args{true}, want: sql.NullBool{Bool: true, Valid: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolenToSqlNullBoolen(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoolenToSqlNullBoolen() = %v, want %v", got, tt.want)
			}
		})
	}
}
