package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	TestNewOrder()
}

func TestNewOrder() {
	apiKey := "FXhZWcpX8TetsYTMk5f9FfrH5j39MUglI6b5pbkYTAsx72KxsErO9eDUpKEE3Ah7"
	secretKey := "IQzutlBhGZNsfzyCgJxlwujpgUHtmOOl80zCnYy9qfditUfkV98wnGgxioWmICqf"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Binance Test New Order endpoint - POST /api/v3/order/test
	testNewOrder, err := client.NewTestNewOrder().Symbol("BTCUSDT").
		Side("BUY").OrderType("MARKET").Quantity(0.001).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(binance_connector.PrettyPrint(testNewOrder))
}
