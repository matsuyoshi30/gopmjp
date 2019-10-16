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

func CalcTenure(name string) int {
	tenure := 0

	for i := 0; i < len(pms)-1; i++ {
		if name == pms[i].Name {
			tenure += countDays(pms[i].StartYear, pms[i].StartMonth, pms[i].StartDay,
				pms[i+1].StartYear, pms[i+1].StartMonth, pms[i+1].StartDay)

			if name != pms[i+1].Name {
				tenure++
			}
		}
	}

	if name == pms[len(pms)-1].Name {
		year, month, day := time.Now().Date()
		tenure += countDays(pms[len(pms)-1].StartYear, pms[len(pms)-1].StartMonth, pms[len(pms)-1].StartDay,
			year, int(month), day)

		tenure++
	}

	return tenure
}

func countDays(sy, sm, sd int, ey, em, ed int) int {
	start := time.Date(sy, time.Month(sm), sd, 0, 0, 0, 0, time.UTC)
	end := time.Date(ey, time.Month(em), ed, 0, 0, 0, 0, time.UTC)
	diff := end.Sub(start)

	h, _ := time.ParseDuration(diff.String())
	return int(h.Hours() / 24.0)
}
