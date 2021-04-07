package util

import "time"

func UnixTime() int64 {
	return time.Now().Unix()
}

func UnixNanTime() int64 {
	return time.Now().UnixNano()
}

func Year() int {
	return time.Now().Year()
}
