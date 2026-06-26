package main

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestDecimalToThaiBaht(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{1234, "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน"},
		{33333.75, "สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์"},
		{0, "ศูนย์บาทถ้วน"},
		{1, "หนึ่งบาทถ้วน"},
		{10, "สิบบาทถ้วน"},
		{11, "สิบเอ็ดบาทถ้วน"},
		{20, "ยี่สิบบาทถ้วน"},
		{21, "ยี่สิบเอ็ดบาทถ้วน"},
		{100, "หนึ่งร้อยบาทถ้วน"},
		{1000, "หนึ่งพันบาทถ้วน"},
		{10000, "หนึ่งหมื่นบาทถ้วน"},
		{100000, "หนึ่งแสนบาทถ้วน"},
		{1000000, "หนึ่งล้านบาทถ้วน"},
		{0.50, "ศูนย์บาทห้าสิบสตางค์"},
		{0.01, "ศูนย์บาทหนึ่งสตางค์"},
		{1000001.01, "หนึ่งล้านหนึ่งบาทหนึ่งสตางค์"},
		{123456789.99, "หนึ่งร้อยยี่สิบสามล้านสี่แสนห้าหมื่นหกพันเจ็ดร้อยแปดสิบเก้าบาทเก้าสิบเก้าสตางค์"},
		{2000000000000, "สองล้านล้านบาทถ้วน"},
		{21000000000000.01, "ยี่สิบเอ็ดล้านล้านบาทหนึ่งสตางค์"},
		// {21000150000000000000.05, "ยี่สิบเอ็ดล้านล้านล้านหนึ่งร้อยห้าสิบล้านล้านบาทห้าสตางค์"}, << this case can cause overflow , the integer part is not 2^53 in float with 64bit register.
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%.2f", tt.input), func(t *testing.T) {
			// Use string input for large numbers to avoid float64 precision loss
			input, _ := decimal.NewFromString(fmt.Sprintf("%.2f", tt.input))
			result := DecimalToThaiBaht(input)
			if result != tt.expected {
				t.Errorf("DecimalToThaiBaht(%v) = %s; want %s", tt.input, result, tt.expected)
			}
		})
	}
}
