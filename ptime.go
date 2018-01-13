package ptime

import (
	"fmt"
	"time"
)

// Unit of Period
type Unit int

const (
	// Day from 00:00 to next day 00:00
	Day Unit = iota

	// Week from monday 00:00 to next monday 00:00
	Week

	// Month from day one of month 00:00 to day one of next month 00:00
	Month

	// Year from day one of year 00:00 to day one of next year 00:00
	Year
)

// Period of time in specific unit,
//
// About timezone: can be any timezone, Start and End should in same timezone.
type Period struct {
	Unit       Unit
	Start, End time.Time
}

// Add n units, returns Period, n can be negative.
func (p Period) Add(n int) Period {
	start, end := getStartEnd(p.Unit, p.Start, n)
	return Period{
		Unit:  p.Unit,
		Start: start,
		End:   end,
	}
}

// New Peroid. parameter t
func New(unit Unit, t time.Time) Period {
	start, end := getStartEnd(unit, t, 0)
	return Period{
		Unit:  unit,
		Start: start,
		End:   end,
	}
}

func getStartEnd(u Unit, t time.Time, n int) (start, end time.Time) {
	y, m, d := t.Date()
	loc := t.Location()
	switch u {
	case Day:
		return time.Date(y, m, d+n, 0, 0, 0, 0, loc),
			time.Date(y, m, d+1+n, 0, 0, 0, 0, loc)

	case Week:
		startd := d - int(t.Weekday()-1) + n*7
		return time.Date(y, m, startd, 0, 0, 0, 0, loc),
			time.Date(y, m, startd+7, 0, 0, 0, 0, loc)

	case Month:
		return time.Date(y, m+time.Month(n), 1, 0, 0, 0, 0, loc),
			time.Date(y, m+time.Month(n)+1, 1, 0, 0, 0, 0, loc)

	case Year:
		return time.Date(y+n, 1, 1, 0, 0, 0, 0, loc),
			time.Date(y+1+n, 1, 1, 0, 0, 0, 0, loc)

	default:
		panic(fmt.Sprintf("Unknown ptime Unit: %d", u))
	}
}
