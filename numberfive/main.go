package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) bool {
	lowerText := strings.ToLower(text)
	tempString := strings.Split(lowerText, "")
	for i := 1; i < len(tempString); i++ {
		start := i
		end := len(tempString) - i - 1
		if tempString[start] != tempString[end] {
			return false
		}
	}
	return true

}
func main() {
	fmt.Println(isPalindrome(("TExT")))
	fmt.Println(isPalindrome(("KodOk")))
	fmt.Println(isPalindrome(("goLang")))
}
