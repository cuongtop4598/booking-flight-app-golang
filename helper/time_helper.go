package helper

import (
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

func ConvertToTime(timestamp *timestamp.Timestamp) *time.Time {
	i, err := strconv.ParseInt(timestamp.String(), 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	return &tm
}

func CaculateDistanceOfTwoTime(t1 time.Time, t2 time.Time) int {
	return int(t2.Sub(t1).Hours())
}

func ConvertTimeToDate(t time.Time) string {
	layoutISO := "2006-01-02"
	return t.Format(layoutISO)
}
