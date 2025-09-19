# Currency Converter CLI Tool (Go)

A command-line tool written in Go that converts one currency to another using live exchange rates from [exchangerate.host](https://exchangerate.host). **Note:** As of now, exchangerate.host requires an API key (free plan available).

---

## Features

- Convert between any two currencies using real-time exchange rates.
- Simple CLI flags for amount, source, and target currency.
- Error handling for invalid inputs or API failures.
- Supports API key input via CLI flag or environment variable.

---

## Installation

1. Clone the repository:
```bash
git clone https://github.com/saitma-op/currency-converter.git
cd currency-converter
```

2. Initialize Go module:
```bash
go mod init currency-converter
```

3. Run the tool (requires API key):
```bash
go run main.go --amount 100 --from USD --to INR --apikey YOUR_KEY
```

---

## Usage

### Convert 100 USD to INR:
```bash
go run main.go --amount 100 --from USD --to INR --apikey YOUR_KEY
```

**Output:**
```
100 USD = 8325.50 INR
```

### CLI Flags

| Flag       | Description                              | Default |
|------------|------------------------------------------|---------|
| `--amount` | Amount to convert                        | 1       |
| `--from`   | Source currency code (e.g., USD)         | USD     |
| `--to`     | Target currency code (e.g., INR)         | INR     |
| `--apikey` | API key for exchangerate.host (required) | none    |

---

## Example Output

```text
50 EUR = 4200.75 INR
200 GBP = 250.10 USD
```

---

## Future Enhancements
- Add `--reverse` flag to swap conversion direction.
- Batch conversions (multiple currency pairs in one run).
- Web interface using Go’s `net/http` or Gin.
- Local caching to reduce API calls.
- Optional fallback to other free APIs like [frankfurter.app](https://www.frankfurter.app).

---

## Contributing

Feel free to fork this repository and submit pull requests with improvements.

---

## License

This project is licensed under the MIT License.