package main

import (
	"fmt"
)

func main() {
	book := "The colour of magic "
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])

	//Slice
	fmt.Println(book[4:11])

	fmt.Println(book[4:])

	fmt.Println(book[:4])

	// concatenate
	fmt.Println("t" + book[1:])

	//multiline
	poem := `
	The road ever on
	Down from the door where it began
	...
	`

	fmt.Println(poem)
}
