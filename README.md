# Thai Baht Currency Converter

A Go service that converts decimal values to Thai text representation with "baht" currency suffix.

## Requirements

- Go 1.16 or higher

## Installation

```bash
# Clone the repository
git clone <your-repo-url>
cd balerion-take-home-assignment

# Install dependencies
go get github.com/shopspring/decimal
```

## Usage

### Run the main program with test cases:

```bash
go run main.go
```

### Build and run:

```bash
go build -o thaibaht
./thaibaht
```

### Run tests:

```bash
go test -v
```

## Examples

| Input | Output |
|-------|--------|
| 1234 | หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน |
| 33333.75 | สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์ |
| 0 | ศูนย์บาทถ้วน |
| 1 | หนึ่งบาทถ้วน |
| 21 | ยี่สิบเอ็ดบาทถ้วน |
| 1000001.01 | หนึ่งล้านหนึ่งบาทหนึ่งสตางค์ |

## API Usage

```go
import (
    "github.com/shopspring/decimal"
)

// Convert decimal to Thai baht text
amount := decimal.NewFromFloat(1234.56)
thaiText := DecimalToThaiBaht(amount)
// Returns: "หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์"
```

## Implementation Notes

- Uses `github.com/shopspring/decimal` for precise decimal arithmetic
- Properly handles Thai number pronunciation rules (e.g., ยี่สิบ for 20, เอ็ด for 1 in certain positions)
- Scales fractional parts by 100 to represent "สตางค์" (satang)
- Returns "ถ้วน" (complete/whole) suffix when there is no fractional part
