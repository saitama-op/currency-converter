package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type ExchangeResponse struct {
	Success bool `json:"success"`
	Query   struct {
		From   string  `json:"from"`
		To     string  `json:"to"`
		Amount float64 `json:"amount"`
	} `json:"query"`
	Info struct {
		Rate float64 `json:"rate"`
	} `json:"info"`
	Result float64 `json:"result"`
}

func main() {
	// CLI flags
	amount := flag.Float64("amount", 1, "Amount to convert")
	from := flag.String("from", "USD", "From currency code")
	to := flag.String("to", "INR", "To currency code")

	flag.Parse()

	url := fmt.Sprintf("https://api.exchangerate.host/convert?from=%s&to=%s&amount=%f", *from, *to, *amount)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching exchange rate:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var data ExchangeResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Error decoding response:", err)
		os.Exit(1)
	}

	if !data.Success {
		fmt.Println("Failed to fetch conversion. Please check input.")
		os.Exit(1)
	}

	fmt.Printf("%.2f %s = %.2f %s\n", *amount, *from, data.Result, *to)
}
