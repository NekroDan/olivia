package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
)

func main() {
	currentTime := time.Now()
	hour := currentTime.Hour()
	activity := ""
	fmt.Println("Hi, Olivia! It is " + currentTime.Format("3:4PM") + ". Press Enter on the keyboard to see what to do!")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	getActivity(hour)
	fmt.Println(activity)
	fmt.Println("You can press Enter again to close this window. Thank you!")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func getActivity(h int) string {
	activity := ""
	switch {
	case h >= 8 && h <= 11:
		fmt.Println("1. Eat breakfast, 2. Try to be quiet if sissy is napping.")
	case h > 11 && h <= 15:
		fmt.Println("1. Eat lunch, 2. Play outside if it's not too hot or rainy.")
	case h > 15 && h <= 18:
		fmt.Println("1. Play with sissy, 2. If mommy starts dinner, go on sissy duty.")
	case h > 19 && h <= 21:
		fmt.Println("1. Start your night time list!")
	default:
		fmt.Println("Ask mommy or daddy.")
	}
	return activity
}

/*func main() {
	currentTime := time.Now()
	hour := time.Now().Hour()
	h,m,s := currentTime.Clock()
	fmt.Println(h,m,s)
	midnight := currentTime.Add( - time.Duration(h)*time.Hour - time.Duration(m) * time.Minute - time.Duration(s)*time.Second)
	fmt.Println(getClock())
func getClock() clock {
	h, m, s := time.Now().Clock()
	return clock(s + 60*(m + 60*h))
}

func (c clock) String() string {

	return "Foo"
}*/