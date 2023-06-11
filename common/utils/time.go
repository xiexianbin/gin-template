package utils

import "time"

func UnixToTime(timestap int) string {
	t := time.Unix(int64(timestap), 0)
	return t.Format("2006-01-02 15:04:05")
}
