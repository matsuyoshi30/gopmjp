package pmjp

import (
	"time"
)

func ToPmjp(year, month, day int) string {
	location, _ := time.LoadLocation("Asia/Tokyo")
	if location == nil {
		location = time.UTC
	}
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 1, location)
	for i := len(pms) - 1; i >= 0; i-- {
		_t := time.Date(pms[i].StartYear, time.Month(pms[i].StartMonth), pms[i].StartDay, 0, 0, 0, 0, location)
		if t.After(_t) {
			return pms[i].Name
		}
	}
	return ""
}

func ToPmjpFromTime(t time.Time) string {
	location, _ := time.LoadLocation("Asia/Tokyo")
	if location == nil {
		location = time.UTC
	}
	for i := len(pms) - 1; i >= 0; i-- {
		_t := time.Date(pms[i].StartYear, time.Month(pms[i].StartMonth), pms[i].StartDay, 0, 0, 0, 0, location)
		if t.After(_t) {
			return pms[i].Name
		}
	}
	return ""
}
