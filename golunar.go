package main


import (
	"fmt"
	"os"
	"path"
	"strconv"

	cc "github.com/lidaobing/chinese_calendar"
  )

func printToday() {
	today := cc.Today()
	fmt.Printf("today:\t%#v\n", today)
}

func printHelp() {
	fmt.Printf("Usage: %s [YEAR MONTH DAY]\n", path.Base(os.Args[0]))
}

func printChineseCalendar2(year, month, day int) (err error) {
	res, err := cc.FromSolarDate(year, month, day)
	if err != nil {
		return err
	}
	fmt.Printf("Result:\t%#v\n", res)
	return nil
}

func printChineseCalendar(year, month, day string) (err error) {
	year2, err := strconv.Atoi(year)
	if err != nil {
		return err
	}

	month2, err := strconv.Atoi(month)
	if err != nil {
		return err
	}

	day2, err := strconv.Atoi(day)
	if err != nil {
		return err
	}

	return printChineseCalendar2(year2, month2, day2)
}

func main() {
	if len(os.Args) == 1 {
		printToday()
		return
	}
	if len(os.Args) != 4 {
		printHelp()
		os.Exit(2)
	}

	err := printChineseCalendar(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
