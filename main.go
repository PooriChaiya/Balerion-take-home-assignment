package main

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

// IDEA:
// 123,456
// 1 -> หนึ่ง-แสน
// 2 -> สอง-หมื่น
// 3 -> สาม-พัน
// 4 -> สี่-ร้อย
// 5 -> ห้า-สิบ **  need to handle หนึ่ง-สิบ and สองสิบ
// 6 -> หก-

// group by 6 digit which will we added ล้าน as suffix

// Implementation :
// 1. split int part and fractional part
// 2. conver int part and fract part -> text
// 3. join with บาท and สตางค์ but need to handle case

var (
	singleDigit = []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
	thUnits = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน"}
)

// DecimalToThaiBaht converts a decimal value to Thai text with baht currency suffix
func DecimalToThaiBaht(d decimal.Decimal) string {
	// Split into integer and fractional parts -> convert to text -> concat with บาท & สตางค์
	intPart := d.Truncate(0)
	fracPart := d.Sub(intPart).Mul(decimal.NewFromInt(100)).Truncate(0)
	intText := convertToThaiText(intPart)

	if fracPart.IsZero() {
		return intText + "บาทถ้วน"
	}

	fracText := convertToThaiText(fracPart)

	// 0.50 if this case call only 50 สตาวค์
	// if intPart.IsZero() {
	// 	return fracText + "สตางค์"
	// }
	return intText + "บาท" + fracText + "สตางค์"
}


//21,000,150,000,000,000,000


func convertToThaiText(d decimal.Decimal) string {
	if d.IsZero() {
		return "ศูนย์"
	}

	numStr := d.String()
	var result strings.Builder

	unitIndex := len(numStr) - 1
	hasValue := false
	countVal := 0

	for i, ch := range numStr {
		digit := int(ch - '0')
		currentUnit := unitIndex - i
		posIdx := currentUnit % 6
		countMil := currentUnit / 6

		
		// need to handle
		// 1. "1" after หลักสิบ change to เอ็ด -> posIdx = 1
		// 2. "2" in the หลักสิบ chnage to ยี่


		if digit == 0 {
			if countVal > 0 && posIdx == 0 && countMil > 0 {
				i := 0 
				for i < countMil {
					result.WriteString("ล้าน")
					i++
				}

				countVal = 0
			}
			continue
		} else {
			countVal++
			hasValue = true
			//1. add number
			if digit == 2 && posIdx == 1 {
				result.WriteString("ยี่")
			} else {
				result.WriteString(singleDigit[digit])
			}
			
			//2. add unit
			// handle หลักสิบ
			result.WriteString(thUnits[posIdx])

			subStr := result.String()
			if strings.HasSuffix(subStr, "หนึ่งสิบ") {
				result.Reset()
				result.WriteString(strings.TrimSuffix(subStr, "หนึ่งสิบ"))
				result.WriteString("สิบ")
			} else if strings.HasSuffix(subStr, "สองสิบ") {
				result.Reset()
				result.WriteString(strings.TrimSuffix(subStr, "สองสิบ"))
				result.WriteString("ยี่สิบ")
			} else if strings.HasSuffix(subStr, "สิบหนึ่ง") {
				result.Reset()
				result.WriteString(strings.TrimSuffix(subStr, "สิบหนึ่ง"))
				result.WriteString("สิบเอ็ด")
			}

		}

		if posIdx == 0 {
			countVal = 0
		}

		// handle million
		if countMil > 0 && posIdx == 0 {
			i := 0 
			for i < countMil {
				result.WriteString("ล้าน")
				i++
			}
		}

		




		
	}

	if !hasValue {
		return singleDigit[0] 
	}

	return result.String()
}

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
		decimal.NewFromFloat(0),
		decimal.NewFromFloat(1),
		decimal.NewFromFloat(10),
		decimal.NewFromFloat(21),
		decimal.NewFromFloat(1000001.01),
		decimal.NewFromFloat(2000000000000),
		decimal.NewFromFloat(21000150000000000000.01),
	}

	fmt.Println("=== Decimal to Thai Baht Converter ===")
	for _, input := range inputs {
		result := DecimalToThaiBaht(input)
		fmt.Printf("Input: %s → Output: %s\n", input.String(), result)
	}
}
