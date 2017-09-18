package main


import (
	"fmt"
	"os"
	"path"
	"strconv"
	"flag"
	"strings"

	cc "github.com/lidaobing/chinese_calendar"
  )

var inverse = flag.Bool("i", false, "from lunar to solar")
var leap = flag.Bool("l", false, "leap month")

func printToday() {
	today := cc.Today()
	fmt.Printf("today:\t%#v\n", today)
}

func printHelp() {
	fmt.Printf("Usage: %s today\n", path.Base(os.Args[0]))
	fmt.Printf("Usage: %s [options] [YEAR MONTH DAY]\n", path.Base(os.Args[0]))
	fmt.Printf("\n")
	fmt.Printf("Options:\n")
	flag.PrintDefaults()
}

func printSolarCalendar(year, month, day int, isLeap bool) (err error) {
	var res cc.ChineseCalendar
	res.Year = year
	res.Month = month
	res.Day = day
	res.IsLeapMonth = isLeap
	t, err := res.ToTime()
	if err != nil {
		return err
	}
	fmt.Printf("Result: %s\n", t)
	return nil
}	

func printChineseCalendar(year, month, day int) (err error) {
	res, err := cc.FromSolarDate(year, month, day)
	if err != nil {
		return err
	}
	fmt.Printf("Result:\t%#v\n", res)
	return nil
}

func str2int(year, month, day string) (year2, month2, day2 int, err error) {
	year2, err = strconv.Atoi(year)
	if err != nil {
		return
	}

	month2, err = strconv.Atoi(month)
	if err != nil {
		return
	}

	day2, err = strconv.Atoi(day)
	if err != nil {
		return
	}
	return
}


func main() {
	flag.Usage = func() {
		printHelp()
	}

	flag.Parse()

	if len(flag.Args()) == 1 && strings.ToUpper(flag.Arg(0)) == "TODAY" {
		printToday()
		return
	}

	if len(flag.Args()) != 3 {
		printHelp()
		os.Exit(2)
	}

	year2, month2, day2, err := str2int(flag.Args()[0], flag.Args()[1], flag.Args()[2])
	
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	if(*inverse) {
		err = printSolarCalendar(year2, month2, day2, *leap)		
	}  else {
		err = printChineseCalendar(year2, month2, day2)
	}
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
