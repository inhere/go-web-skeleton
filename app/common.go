package app

import (
	"time"
)

// FormatPageAndSize
func FormatPageAndSize(page int, size int) (int, int) {
	if page < 1 {
		page = 1
	}

	if size > MaxPageSize {
		size = MaxPageSize
	}

	return page, size
}

// LocUnixTime get local unix time
func LocUnixTime() int64 {
	return time.Now().Local().Unix()
}

// PRCTime get PRC local time
func PRCTime() time.Time {
	loc, _ := time.LoadLocation(Timezone)

	return time.Now().In(loc)
}

// AbsPath always return abs path.
func AbsPath(path string) string {
	if string(path[0]) == "/" {
		return path
	}

	return path
}
