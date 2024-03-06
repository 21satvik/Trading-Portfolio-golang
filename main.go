package main

import (
    "log"
    "net/http"
)

var (
    portfolio      Portfolio
    shouldStop     = false        // Flag to indicate whether WebSocket should stop
    stop           = make(chan bool, 1) // Buffered stop channel
)

func main() {
    // Start WebSocket connection to Binance
    go connectToBinance(&portfolio, stop)

    // Define HTTP routes
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", homeHandler)

    http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
        AddCoinHandler(w, r, &portfolio)
    })
    http.HandleFunc("/remove", func(w http.ResponseWriter, r *http.Request) {
        RemoveCoinHandler(w, r, &portfolio)
    })
    http.HandleFunc("/portfolio", func(w http.ResponseWriter, r *http.Request) {
        PortfolioHandler(w, r, &portfolio)
    })

    // Start HTTP server
    log.Println("Server listening on localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
