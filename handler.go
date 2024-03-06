package main

import (
    "encoding/json"
    "net/http"
)

// AddCoinHandler handles adding a new coin to the portfolio
func AddCoinHandler(w http.ResponseWriter, r *http.Request, portfolio *Portfolio) {
    var newCoin Coinset
    err := json.NewDecoder(r.Body).Decode(&newCoin)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
    }

    // Add the coin to the portfolio (implementation omitted)
    portfolio.AddCoin(newCoin)
    w.WriteHeader(http.StatusOK)

    // Stop WebSocket goroutine and restart with updated portfolio
    stop <- true
    go connectToBinance(portfolio, stop)

}

// RemoveCoinHandler handles removing a coin from the portfolio
func RemoveCoinHandler(w http.ResponseWriter, r *http.Request, portfolio *Portfolio) {
    var requestData struct {
        CoinName string  `json:"coin"`
        Amount   float64 `json:"amount"`
    }
    err := json.NewDecoder(r.Body).Decode(&requestData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Remove the coin from the portfolio (implementation omitted)
    portfolio.RemoveCoin(requestData.CoinName, requestData.Amount)
    w.WriteHeader(http.StatusOK)

    // Stop WebSocket goroutine and restart with updated portfolio
    stop <- true
    go connectToBinance(portfolio, stop)

}

// PortfolioHandler handles showing the portfolio
func PortfolioHandler(w http.ResponseWriter, r *http.Request, portfolio *Portfolio) {
    jsonData, err := json.Marshal(portfolio)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonData)
}
