package pmjp

import (
	"time"
)

func ToPmjp(t time.Time) string {
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
