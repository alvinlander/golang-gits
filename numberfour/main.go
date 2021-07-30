package main

import (
	"fmt"
	"strings"
)

func revString(text string) {
	var result []string
	tempString := strings.Split(text, "")
	for i := 0; i < len(tempString); i++ {
		result = append(result, tempString[len(tempString)-i-1])
	}
	fmt.Println(strings.Join(result, ""))

}
func main() {
	revString("Halo")
	revString("Aplikasi")
	revString("isakilpA")
}
