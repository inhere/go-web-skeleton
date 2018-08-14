package app

import "time"

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

// LocTime get local time
func LocTime() time.Time {
	// loc, _ := time.LoadLocation(Timezone)
	// return time.Now().In(loc)
	return time.Now().Local()
}

// PRCTime get PRC local time
func PRCTime() time.Time {
	loc, _ := time.LoadLocation(Timezone)

	return time.Now().In(loc)
}
