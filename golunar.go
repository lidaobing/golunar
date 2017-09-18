package main


import (
	"fmt"

	cc "github.com/lidaobing/chinese_calendar"
  )

func main() {
	today := cc.Today()
	fmt.Printf("today:\t%#v\n", today)
}
