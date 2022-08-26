package timeDate

// https://golangbyexample.com/all-about-time-and-date-golang/
import (
	"fmt"
	"time"
)

/*
type Time struct {
	wall uint64
	ext  int64

	loc *Location
}
type Duration int64
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)
func Now() Time
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
func (t Time) Add(d Duration) Time
func (t Time) AddDate(years int, months int, days int) Time
*/

func TimePrint() {
	t := time.Now()

	//Add 1 hours
	newT_1h := t.Add(time.Hour * 1)
	fmt.Printf("Adding 1 hour\n: %s\n", newT_1h)

	//Add 15 min
	newT_15m := t.Add(time.Minute * 15)
	fmt.Printf("Adding 15 minute\n: %s\n", newT_15m)

	//Add 10 sec
	newT_10s := t.Add(time.Second * 10)
	fmt.Printf("Adding 10 sec\n: %s\n", newT_10s)

	//Add 100 millisecond
	newT_100ms := t.Add(time.Millisecond * 10)
	fmt.Printf("Adding 100 millisecond\n: %s\n", newT_100ms)

	//Add 1000 microsecond
	newT_1000ms := t.Add(time.Millisecond * 10)
	fmt.Printf("Adding 1000 microsecond\n: %s\n", newT_1000ms)

	//Add 10000 nanosecond
	newT_10000ns := t.Add(time.Nanosecond * 10000)
	fmt.Printf("Adding 1000 nanosecond\n: %s\n", newT_10000ns)

	//Add 1 year 2 month 4 day
	newT_etc := t.AddDate(1, 2, 4)
	fmt.Printf("Adding 1 year 2 month 4 day\n: %s\n", newT_etc)

	/*
		Adding 1 hour:
		2020-02-01 02:16:35.893847 +0530 IST m=+3600.000239893

		Adding 15 minute:
		2020-02-01 01:31:35.893847 +0530 IST m=+900.000239893

		Adding 10 sec:
		2020-02-01 01:16:45.893847 +0530 IST m=+10.000239893

		Adding 100 millisecond:
		2020-02-01 01:16:35.903847 +0530 IST m=+0.010239893

		Adding 1000 microsecond:
		2020-02-01 01:16:35.903847 +0530 IST m=+0.010239893

		Adding 1000 nanosecond:
		2020-02-01 01:16:35.893857 +0530 IST m=+0.000249893

		Adding 1 year 2 month 4 day:
		2021-04-05 01:16:35.893847 +0530 IST
	*/

	//Add 1 hours
	newT_s1h := t.Add(-time.Hour * 1)
	fmt.Printf("Subtracting 1 hour:\n %s\n", newT_s1h)

	//Add 15 min
	newT_s15m := t.Add(-time.Minute * 15)
	fmt.Printf("Subtracting 15 minute:\n %s\n", newT_s15m)

	//Add 10 sec
	newT_s10s := t.Add(-time.Second * 10)
	fmt.Printf("Subtracting 10 sec:\n %s\n", newT_s10s)

	//Add 100 millisecond
	newT_s100ms := t.Add(-time.Millisecond * 10)
	fmt.Printf("Subtracting 100 millisecond:\n %s\n", newT_s100ms)

	//Add 1000 microsecond
	newT_s1000ms := t.Add(-time.Millisecond * 10)
	fmt.Printf("Subtracting 1000 microsecond:\n %s\n", newT_s1000ms)

	//Add 10000 nanosecond
	newT_s10000ns := t.Add(-time.Nanosecond * 10000)
	fmt.Printf("Subtracting 1000 nanosecond:\n %s\n", newT_s10000ns)

	//Add 1 year 2 month 4 day
	newT_sEtc := t.AddDate(-1, -2, -4)
	fmt.Printf("Subtracting 1 year 2 month 4 day:\n %s\n", newT_sEtc)
	/*
	   Subtracting 1 hour:
	    2020-02-01 00:18:29.772673 +0530 IST m=-3599.999784391

	   Subtracting 15 minute:
	    2020-02-01 01:03:29.772673 +0530 IST m=-899.999784391

	   Subtracting 10 sec:
	    2020-02-01 01:18:19.772673 +0530 IST m=-9.999784391

	   Subtracting 100 millisecond:
	    2020-02-01 01:18:29.762673 +0530 IST m=-0.009784391

	   Subtracting 1000 microsecond:
	    2020-02-01 01:18:29.762673 +0530 IST m=-0.009784391

	   Subtracting 1000 nanosecond:
	    2020-02-01 01:18:29.772663 +0530 IST m=+0.000205609

	   Subtracting 1 year 2 month 4 day:
	    2018-11-27 01:18:29.772673 +0530 IST
	*/

}
