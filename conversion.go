package goutils

import (
	"database/sql"
	"strconv"
	"strings"
	"time"
)

func StringToInt64(key string) (int64, error) {
	return strconv.ParseInt(key, 10, 64)
}

func StringToBoolen(key string) (bool, error) {
	return strconv.ParseBool(key)
}

func StringToInt(key string) (int, error) {
	return strconv.Atoi(key)
}

func StringToFloat64(key string) (float64, error) {
	return strconv.ParseFloat(key, 64)
}

func Int64ToString(value int64) string {
	return strconv.FormatInt(value, 10)
}

func Float64ToString(key float64) string {
	return strconv.FormatFloat(key, 'f', -1, 64)
}

func SqlNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func TimeToSqlNullTime(t time.Time) sql.NullTime {
	if strings.Contains(t.String(), "1970-01-01") || strings.Contains(t.String(), "0001-01-01") {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

func Int64ToSqlNullInt64(t int64) sql.NullInt64 {
	if t == 0 {
		return sql.NullInt64{}
	}

	return sql.NullInt64{
		Int64: t,
		Valid: true,
	}
}

func Float64ToSqlNullFloat64(t float64) sql.NullFloat64 {
	if t == 0 {
		return sql.NullFloat64{}
	}

	return sql.NullFloat64{
		Float64: t,
		Valid:   true,
	}
}

func BoolenToSqlNullBoolen(t bool) sql.NullBool {
	if t == false {
		return sql.NullBool{}
	}

	return sql.NullBool{
		Bool:  t,
		Valid: true,
	}
}
