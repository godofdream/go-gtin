package gtin

import (
	"math"
)

// inspired by https://github.com/TheGeekExplorer/GTIN-Validation/blob/master/GO/main.go

//pseudo-constant
var gtinMaths = [17]int{3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3}

// IsGTIN returns true if given string is a gtin and returns false otherwise
func IsGTIN(gtin string) bool {
	var gtinLength = len(gtin)             // Length of GTIN
	if gtinLength < 8 || gtinLength > 14 { //Check if length is ok for GTIN
		return false
	}

	var CheckDigitArray = make([]int, 17, 17) // Array of calculations
	var modifier = 18 - gtinLength            // Modifier for Calc Table (17 - (gtinLength - 1))

	// Split string at chars into an array
	var BarcodeArray = make([]int, 17, 17)
	for i, r := range gtin {
		BarcodeArray[i] = int(r)
	}

	// Run through and put digits into multiplication table
	for i := 0; i < (gtinLength - 1); i++ {
		barcodeDigit := int(gtin[i] - '0')
		if barcodeDigit < 0 || barcodeDigit > 9 { // check if digit is a number
			return false
		}
		CheckDigitArray[modifier+i] = barcodeDigit // Add barcode digit to Multiplication Table
	}

	// Temp Variables
	var tmpCheckSum = 0
	// Calculate "Sum" of barcode digits
	for i := modifier; i < 17; i++ {
		tmpCheckSum = tmpCheckSum + (CheckDigitArray[i] * gtinMaths[i])
	}
	tmpCheckDigit := float64(tmpCheckSum)
	tmpCheckDigit = (math.Ceil(tmpCheckDigit/10) * 10) - tmpCheckDigit

	// compare CheckDigit with calculated one
	if int(gtin[gtinLength-1]-'0') == int(tmpCheckDigit) {
		return true
	}
	return false
}
