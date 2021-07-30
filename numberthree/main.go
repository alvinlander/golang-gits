package main

import (
	"fmt"
	"strconv"
)

func timeConversion(time string) string {
	if time[:2] == "12" && time[len(time)-2:] == "AM" {
		return "00" + time[2:5]
	} else if time[:2] == "12" && time[len(time)-2:] == "PM" {
		return time[:5]
	} else if time[len(time)-2:] == "AM" {
		return time[:5]
	} else {
		firstDigit, _ := strconv.Atoi(time[:2])
		convFristDigit := strconv.Itoa(firstDigit)
		return "0" + convFristDigit + time[2:5]
	}
}
func main() {
	fmt.Println(timeConversion("12:00:00 AM"))
	fmt.Println(timeConversion("12:00:00 PM"))
	fmt.Println(timeConversion("08:08:00 AM"))
	fmt.Println(timeConversion("09:00:00 PM"))
}
