package main

// Coinset represents a cryptocurrency coin
type Coinset struct {
	Coin         string  `json:"coin"`
	PurchasePrice float64 `json:"price"`
	CurrentPrice  float64 `json:"current"`
	Amount       float64 `json:"amount"`
	PLamount   float64 `json:"plamount"`
	PLpercent   float64 `json:"plpercent"`
}

// Portfolio represents a collection of coins
type Portfolio struct {
	Portfolio []Coinset `json:"portfolio"`
	TotalPLamount float64 `json:"totalPLamount"`
	TotalPLpercent float64 `json:"totalPLpercent"`
}

// AddCoin adds a new coin to the portfolio or updates existing coin if it already exists
func (p *Portfolio) AddCoin(newCoin Coinset) {
	// Check if the coin already exists in the portfolio
	for i, coin := range p.Portfolio {
		if coin.Coin == newCoin.Coin {
			// Update existing coin by averaging price and adding amounts
			totalAmount := coin.Amount + newCoin.Amount
			totalPurchasePrice := coin.PurchasePrice*coin.Amount + newCoin.PurchasePrice*newCoin.Amount
			p.Portfolio[i].PurchasePrice = totalPurchasePrice / totalAmount
			p.Portfolio[i].Amount = totalAmount
			return
		}
	}

	// If the coin doesn't exist, add it to the portfolio
	p.Portfolio = append(p.Portfolio, newCoin)
}

// RemoveCoin removes a specified amount of coins from the portfolio
func (p *Portfolio) RemoveCoin(coinName string, amountToRemove float64) {
	for i, coin := range p.Portfolio {
		if coin.Coin == coinName {
			// Check if the amount to remove is greater than or equal to the current amount
			if amountToRemove >= coin.Amount {
				// Remove the entire coin entry from the portfolio
				p.Portfolio = append(p.Portfolio[:i], p.Portfolio[i+1:]...)
				return
			} else {
				// Update the amount of the coin in the portfolio
				p.Portfolio[i].Amount -= amountToRemove
				return
			}
		}
	}
}

// CalculateProfitLoss calculates the profit/loss for each coin and the total profit/loss for the portfolio
func (p *Portfolio) CalculateProfitLoss() {
	var (
		TotalPurchaseAmount float64
		TotalCurrentAmount float64
		TotalPLamount float64
	)

	// Calculate profit/loss for each coin
	for i, coin := range p.Portfolio {
		TotalPurchaseAmount += coin.PurchasePrice * coin.Amount
		TotalCurrentAmount += coin.CurrentPrice * coin.Amount
		p.Portfolio[i].PLamount = (coin.CurrentPrice - coin.PurchasePrice) * coin.Amount
		p.Portfolio[i].PLpercent = ((coin.CurrentPrice - coin.PurchasePrice) / coin.PurchasePrice) * 100
		TotalPLamount += p.Portfolio[i].PLamount
	}

	// Calculate total profit/loss for the portfolio
	p.TotalPLamount = TotalPLamount
	p.TotalPLpercent = ((TotalCurrentAmount - TotalPurchaseAmount) / TotalPurchaseAmount) * 100
}