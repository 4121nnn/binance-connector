package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetAllOrders()
}

func GetAllOrders() {
	apiKey := "FXhZWcpX8TetsYTMk5f9FfrH5j39MUglI6b5pbkYTAsx72KxsErO9eDUpKEE3Ah7"
	secretKey := "IQzutlBhGZNsfzyCgJxlwujpgUHtmOOl80zCnYy9qfditUfkV98wnGgxioWmICqf"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Binance Get all account orders; active, canceled, or filled - GET /api/v3/allOrders
	getAllOrders, err := client.NewGetAllOrdersService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getAllOrders))
}
