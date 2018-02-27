package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello time")
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2017, 2, 27, 20, 30, 10, 2131231, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))

	secs := now.Unix()
	nanos := now.UnixNano()
	millis := nanos / 1000000

	p(secs)
	p(millis)
	p(nanos)

	p(time.Unix(secs, 0))
	p(time.Unix(0, nanos))
}
