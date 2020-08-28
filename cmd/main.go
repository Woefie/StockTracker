package main

import (
	"context"
	"fmt"

	"github.com/Finnhub-Stock-API/finnhub-go"
)

func main() {

	finnhubClient := finnhub.NewAPIClient(finnhub.NewConfiguration()).DefaultApi
	auth := context.WithValue(context.Background(), finnhub.ContextAPIKey, finnhub.APIKey{
		Key: "bt4fk1f48v6u8ohnujl0", // Replace this
	})

	//Stock candles
	stockCandles, _, err := finnhubClient.StockCandles(auth, "AAPL", "D", 1590988249, 1591852249, nil)
	if(err != nil){
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", stockCandles)

	
}
