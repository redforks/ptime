package ptime_test

import (
	"time"

	. "github.com/redforks/ptime"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ptime", func() {
	var (
		aTime  = time.Date(2018, 2, 13, 16, 45, 0, 0, time.Local)
		loc2   = time.FixedZone("foo", -6)
		aTime2 = time.Date(2018, 2, 13, 16, 45, 0, 0, loc2)

		newDateTZ = func(year, month, day int, loc *time.Location) time.Time {
			return time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)
		}

		newDate = func(year, month, day int) time.Time {
			return newDateTZ(year, month, day, time.Local)
		}

		newDate2 = func(year, month, day int) time.Time {
			return newDateTZ(year, month, day, loc2)
		}
	)

	DescribeTable("New", func(u Unit, t time.Time, expStart, expEnd time.Time) {
		p := New(u, t)
		Ω(p.Unit).Should(Equal(u))
		Ω(p.Start).Should(Equal(expStart))
		Ω(p.End).Should(Equal(expEnd))
		Ω(p.Start.Location()).Should(Equal(t.Location()))
	},
		Entry("Day", Day, aTime, newDate(2018, 2, 13), newDate(2018, 2, 14)),
		Entry("Day in alter TZ", Day, aTime2, newDate2(2018, 2, 13), newDate2(2018, 2, 14)),
		Entry("Week", Week, aTime, newDate(2018, 2, 12), newDate(2018, 2, 19)),
		Entry("Month", Month, aTime, newDate(2018, 2, 1), newDate(2018, 3, 1)),
		Entry("Year", Year, aTime, newDate(2018, 1, 1), newDate(2019, 1, 1)),
	)

	DescribeTable("Add", func(u Unit, t time.Time, n int, expStart, expEnd time.Time) {
		p := New(u, t).Add(n)
		Ω(p.Unit).Should(Equal(u))
		Ω(p.Start).Should(Equal(expStart))
		Ω(p.End).Should(Equal(expEnd))
		Ω(p.Start.Location()).Should(Equal(t.Location()))
	},
		Entry("Day", Day, aTime, 3, newDate(2018, 2, 16), newDate(2018, 2, 17)),
		Entry("Prev Day", Day, aTime, -2, newDate(2018, 2, 11), newDate(2018, 2, 12)),
		Entry("Week", Week, aTime, 1, newDate(2018, 2, 19), newDate(2018, 2, 26)),
		Entry("Month", Month, aTime, -2, newDate(2017, 12, 1), newDate(2018, 1, 1)),
		Entry("Year", Year, aTime, 1, newDate(2019, 1, 1), newDate(2020, 1, 1)),
	)
})
