package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61, //Must have trailing comma in multi line
	}

	//number of items
	fmt.Println(len(stocks))

	//get a value
	fmt.Println(stocks["MSFT"])

	//get zero value if not found
	fmt.Println(stocks["TSLA"])

	// use two value to see if found
	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	// set
	stocks["TSLA"] = 322.12
	fmt.Println(stocks)

	//delete
	delete(stocks, "AMZN")
	fmt.Println(stocks)

	//single value for is on keys
	for key := range stocks {
		fmt.Println(key)
	}

	//double value for is key , value
	for key, value := range stocks {
		fmt.Printf("%s --> %.2f\n", key, value)
	}

}
