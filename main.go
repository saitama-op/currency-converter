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
	Error  *struct {
		Code int    `json:"code"`
		Type string `json:"type"`
		Info string `json:"info"`
	} `json:"error,omitempty"`
}

func main() {
	amount := flag.Float64("amount", 1, "Amount to convert")
	from := flag.String("from", "USD", "Source currency code")
	to := flag.String("to", "INR", "Target currency code")
	apiKey := flag.String("apikey", "", "API key for exchangerate.host")
	flag.Parse()

	if *apiKey == "" {
		fmt.Println("Error: API key is required. Use --apikey=YOUR_KEY")
		os.Exit(1)
	}

	url := fmt.Sprintf(
		"https://api.exchangerate.host/convert?from=%s&to=%s&amount=%f&access_key=%s",
		*from, *to, *amount, *apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching rate:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var data ExchangeResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Error decoding response:", err)
		os.Exit(1)
	}

	if !data.Success {
		if data.Error != nil {
			fmt.Printf("API error: %s - %s\n", data.Error.Type, data.Error.Info)
		} else {
			fmt.Println("Failed to fetch rate.")
		}
		os.Exit(1)
	}

	fmt.Printf("%.2f %s = %.2f %s\n", *amount, *from, data.Result, *to)
}
