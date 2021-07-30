package main

import "fmt"

func helloWorld(num int) {
	if (num%3 == 0) && (num%5 == 0) {
		fmt.Println("Hello World")
	} else if num%3 == 0 {
		fmt.Println("Hello")
	} else if num%5 == 0 {
		fmt.Println("World")
	}
}
func main() {
	helloWorld(15)
	helloWorld(10)
	helloWorld(9)
}
