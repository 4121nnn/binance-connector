package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetMyTrades()
}

func GetMyTrades() {
	apiKey := "FXhZWcpX8TetsYTMk5f9FfrH5j39MUglI6b5pbkYTAsx72KxsErO9eDUpKEE3Ah7"
	secretKey := "IQzutlBhGZNsfzyCgJxlwujpgUHtmOOl80zCnYy9qfditUfkV98wnGgxioWmICqf"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Binance Get trades for a specific account and symbol (USER_DATA) - GET /api/v3/myTrades
	getMyTradesService, err := client.NewGetMyTradesService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getMyTradesService))
}
