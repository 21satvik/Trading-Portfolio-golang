function fetchPortfolioData() {
    fetch('/portfolio')
        .then(response => response.json())
        .then(data => {
            const portfolioTable = document.getElementById('portfolioTable');
            const totalPLamount = document.getElementById('totalPLamount');
            const totalPLpercent = document.getElementById('totalPLpercent');

            // Clear existing table rows
            portfolioTable.innerHTML = '';
            if (data.portfolio) {
                // Populate table with updated portfolio data
                data.portfolio.forEach(coin => {
                    const row = document.createElement('tr');
                    row.innerHTML = `
                            <td>${coin.coin}</td>
                            <td>${coin.price.toFixed(2)}</td>
                            <td>${coin.current.toFixed(2)}</td>
                            <td>${coin.amount}</td>
                            <td>${coin.plamount.toFixed(2)}</td>
                            <td>${coin.plpercent.toFixed(2)}</td>
                        `;
                    portfolioTable.appendChild(row);
                });

                // Update total profit/loss
                totalPLamount.textContent = data.totalPLamount.toFixed(2);
                totalPLpercent.textContent = data.totalPLpercent.toFixed(2);
            }

        })
        .catch(error => console.error('Error fetching portfolio data:', error));
}

// Fetch portfolio data initially when the page loads
fetchPortfolioData();

// Fetch portfolio data every 10 seconds (adjust interval as needed)
setInterval(fetchPortfolioData, 1000); // 1 seconds

// Function to handle form submission
document.getElementById('addCoinForm').addEventListener('submit', function (event) {
    event.preventDefault(); // Prevent default form submission

    // Get input values
    const coin = document.getElementById('coin').value;
    const price = parseFloat(document.getElementById('price').value);
    const amount = parseFloat(document.getElementById('amount').value);

    // Validate input values
    if (isNaN(price) || isNaN(amount)) {
        alert('Please enter valid numeric values for price and amount.');
        return;
    }

    // Send a POST request to the server to add the new coin
    fetch('/add', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ coin, price, amount })
    })
        .then(response => response.json())
        .then(data => {
            // Handle response from server
            console.log('New coin added:', data);
        })
        .catch(error => {
            console.error('Error adding coin:', error);
        });
});

// Function to handle removing a coin from the portfolio
document.getElementById('removeCoinForm').addEventListener('submit', function (event) {
    event.preventDefault(); // Prevent default form submission

    // Get input value
    const coin = document.getElementById('removeCoin').value;
    const amount = parseFloat(document.getElementById('removeAmount').value);

    // Validate input value
    if (isNaN(amount)) {
        alert('Please enter a valid numeric value for amount.');
        return;
    }

    // Send a POST request to the server to remove the coin
    fetch('/remove', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ coin, amount })
    })
        .then(response => response.json())
        .then(data => {
            // Handle response from server
            console.log('Coin removed:', data);
        })
        .catch(error => {
            console.error('Error removing coin:', error);
        });
});