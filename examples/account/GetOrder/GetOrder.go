package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	QueryOrder()
}

func QueryOrder() {
	apiKey := "FXhZWcpX8TetsYTMk5f9FfrH5j39MUglI6b5pbkYTAsx72KxsErO9eDUpKEE3Ah7"
	secretKey := "IQzutlBhGZNsfzyCgJxlwujpgUHtmOOl80zCnYy9qfditUfkV98wnGgxioWmICqf"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Binance Query Order (USER_DATA) - GET /api/v3/order
	queryOrder, err := client.NewGetOrderService().Symbol("BTCUSDT").
		OrderId(20064739).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(queryOrder))
}
