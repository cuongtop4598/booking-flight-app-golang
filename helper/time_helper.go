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
