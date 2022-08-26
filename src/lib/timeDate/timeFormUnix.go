package timeDate

// https://golangbyexample.com/all-about-time-and-date-golang/
import (
	"fmt"
	"time"
)

func TimePrint2() {

	//Parse YYYY-MM-DD
	timeT, _ := time.Parse("2006-01-02", "2020-01-29")
	fmt.Println(timeT)

	//Parse YY-MM-DD
	timeT, _ = time.Parse("06-01-02", "20-01-29")
	fmt.Println(timeT)

	//Parse YYYY-#{MonthName}-DD
	timeT, _ = time.Parse("2006-Jan-02", "2020-Jan-29")
	fmt.Println(timeT)

	//Parse YYYY-#{MonthName}-DD WeekDay HH:MM:SS
	timeT, _ = time.Parse("2006-Jan-02 Monday 03:04:05", "2020-Jan-29 Wednesday 12:19:25")
	fmt.Println(timeT)

	//Parse YYYY-#{MonthName}-DD WeekDay HH:MM:SS PM Timezone TimezoneOffset
	timeT, _ = time.Parse("2006-Jan-02 Monday 03:04:05 PM MST -07:00", "2020-Jan-29 Wednesday 12:19:25 AM IST +05:30")
	fmt.Println(timeT)
	/*
	   2020-01-29 00:00:00 +0000 UTC
	   2020-01-29 00:00:00 +0000 UTC
	   2020-01-29 00:00:00 +0000 UTC
	   2020-01-29 12:19:25 +0000 UTC
	   2020-01-29 00:19:25 +0530 ISTS
	*/

	now := time.Now()

	//Format YYYY-MM-DD
	fmt.Printf("YYYY-MM-DD: %s\n", now.Format("2006-01-02"))

	//Format YY-MM-DD
	fmt.Printf("YY-MM-DD: %s\n", now.Format("06-01-02"))

	//Format YYYY-#{MonthName}-DD
	fmt.Printf("YYYY-#{MonthName}-DD: %s\n", now.Format("2006-Jan-02"))

	//Format HH:MM:SS
	fmt.Printf("HH:MM:SS: %s\n", now.Format("03:04:05"))

	//Format HH:MM:SS Millisecond
	fmt.Printf("HH:MM:SS Millisecond: %s\n", now.Format("03:04:05 .999"))

	//Format YYYY-#{MonthName}-DD WeekDay HH:MM:SS
	fmt.Printf("YYYY-#{MonthName}-DD WeekDay HH:MM:SS: %s\n", now.Format("2006-Jan-02 Monday 03:04:05"))

	//Format YYYY-#{MonthName}-DD WeekDay HH:MM:SS PM Timezone TimezoneOffset
	fmt.Printf("YYYY-#{MonthName}-DD WeekDay HH:MM:SS PM Timezone TimezoneOffset: %s\n", now.Format("2006-Jan-02 Monday 03:04:05 PM MST -07:00"))

	/*
		YYYY-MM-DD: 2020-01-25
		YY-MM-DD: 20-01-25
		YYYY-#{MonthName}-DD: 2020-Jan-25
		HH:MM:SS: 11:14:16
		HH:MM:SS Millisecond: 11:14:16 .213
		YYYY-#{MonthName}-DD WeekDay HH:MM:SS: 2020-Jan-25 Saturday 11:14:16
		YYYY-#{MonthName}-DD WeekDay HH:MM:SS PM Timezone TimezoneOffset: 2020-Jan-25 Saturday 11:14:16 PM IST +05:30
	*/

	//time.Time to Unix Timestamp
	// 	now := time.Now()
	tUnix := now.Unix()
	fmt.Printf("timeUnix %d\n", tUnix)

	//Unix Timestamp to time.Time
	timeT_u := time.Unix(tUnix, 0)
	fmt.Printf("time.Time: %s\n", timeT_u)

	// timeUnix 1257894000
	// time.Time: 2009-11-10 23:00:00 +0000 UTC

	// 	now := time.Now()
	loc, _ := time.LoadLocation("UTC")
	fmt.Printf("UTC Time: %s\n", now.In(loc))

	loc, _ = time.LoadLocation("Europe/Berlin")
	fmt.Printf("Berlin Time: %s\n", now.In(loc))

	loc, _ = time.LoadLocation("America/New_York")
	fmt.Printf("New York Time: %s\n", now.In(loc))

	loc, _ = time.LoadLocation("Asia/Dubai")
	fmt.Printf("Dubai Time: %s\n", now.In(loc))

	/*
		UTC Time: 2020-01-31 18:09:41.705858 +0000 UTC
		Berlin Time: 2020-01-31 19:09:41.705858 +0100 CET
		New York Time: 2020-01-31 13:09:41.705858 -0500 EST
		Dubai Time: 2020-01-31 22:09:41.705858 +0400 +04
	*/

}
