package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			println("fizz buzz")
		} else if i%3 == 0 {
			println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
