package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello World , test!")

		time.Sleep(time.Second * 2)
	}
}
