package main

import (
    "encoding/json"
    "github.com/gorilla/websocket"
    "strings"
	"log"
	"strconv"
	"time"
)

// Function to establish WebSocket connection to Binance API and receive ticker data
func connectToBinance(portfolio *Portfolio, stop chan bool) {
	loop:=true
	// Stop the WebSocket connection when the stop channel receives a message
	go func() {
		<-stop
		log.Println("Stopping WebSocket connection")
		loop=false
		return
	}()

	// Prepare asset strings for WebSocket stream
	var assetStrings []string
	
	for _, coin := range portfolio.Portfolio {
		assetStrings = append(assetStrings, strings.ToLower(coin.Coin)+"@miniTicker")
	}
	assetsStr := strings.Join(assetStrings, "/")
	log.Println("Connecting to Binance WebSocket")

	// Construct WebSocket URL
	socketURL := "wss://stream.binance.com:9443/stream?streams=" + assetsStr
	log.Printf("WebSocket URL: %s", socketURL)

	// Establish WebSocket connection
	conn, _, err := websocket.DefaultDialer.Dial(socketURL, nil)
	if err != nil {
		log.Printf("WebSocket connection error: %v", err)
		return
	}
	defer conn.Close()

	log.Println("WebSocket connection established")

	// Read and process incoming messages
	for {
		if loop==false{
			break
		}
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			return
		}
		
		
		// Unmarshal message into a map
		var data map[string]interface{}
		if err := json.Unmarshal(message, &data); err != nil {
			log.Printf("WebSocket unmarshal error: %v", err)
			continue
		}
		
		// Extract symbol and price from the map
		symbol, ok := data["data"].(map[string]interface{})["s"].(string)
		if !ok {
			log.Println("Symbol not found in message")
			continue
		}
		
		priceStr, ok := data["data"].(map[string]interface{})["c"].(string)
		if !ok {
			log.Println("Price not found in message")
			continue
		}
		
		// Convert price to float
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Printf("Price conversion error: %v", err)
			continue
		}
		
		// Update current price in portfolio
		updateCurrentPrice(portfolio, symbol, price)
		// Calculate profit/loss for the portfolio
		portfolio.CalculateProfitLoss()
		time.Sleep(1 * time.Second)
	}
}

// Function to update current price of a coin in the portfolio
func updateCurrentPrice(portfolio *Portfolio, symbol string, price float64) {
	for i, coin := range portfolio.Portfolio {
		if coin.Coin == symbol {
			portfolio.Portfolio[i].CurrentPrice = price
			return
		}
	}
}