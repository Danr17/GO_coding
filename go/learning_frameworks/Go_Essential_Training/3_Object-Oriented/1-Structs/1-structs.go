package main

import (
	"fmt"
)

//Trade is a trade in stcoks
type Trade struct {
	Symbol string  //Stock symbol
	Volume int     //number of shares
	Price  float64 // trade price
	Buy    bool    // true if buy trade , false if sale trade
}

func main() {
	t1 := Trade{"MSFT", 10, 99.98, true}
	fmt.Println(t1)

	fmt.Printf("%+v\n", t1)

	t2 := Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.98,
		Buy:    true,
	}
	fmt.Printf("%+v\n", t2)

	t3 := Trade{}
	fmt.Printf("%+v\n", t3)
}
