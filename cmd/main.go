package main

import (
	"fmt"

	"github.com/m1/go-finnhub/client"
)

func main() {

	c := client.New("bt4fk1f48v6u8ohnujl0")

	// Stocks
	company, err := c.Stock.GetProfile("AAPL")

	if err != nil {
		fmt.Println(err)
	}

	// Success!
	fmt.Println(company)
}
