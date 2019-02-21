package main

import (
	"fmt"
)

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

//value returns the trade value
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}

func main() {
	t := Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.98,
		Buy:    true,
	}
	fmt.Println(t.Value())
}
