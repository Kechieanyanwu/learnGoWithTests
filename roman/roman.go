package roman

import (
	"log"
	"strings"
)

// new implementation decouples logic from the data

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder
	if arabic > 3999 {
		log.Fatal("You can't convert a number greater than 3999")
	}

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()

}

func ConvertToArabic(roman string) (total uint16) {
	for _, symbols := range windowedRoman(roman).Symbols() {
		total += allRomanNumerals.ValueOf(symbols...)
	}
	return
}

func isSubtractive(currentSymbol uint8) bool {
	return currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C'
}

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
	for _, s := range r {
		if s.Symbol == string(symbols) {
			return s.Value
		}
	}
	return 0
}

func (r RomanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}
	return false
}

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

type windowedRoman string

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)

		if notAtEnd && isSubtractive(symbol) && allRomanNumerals.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{symbol, w[i+1]})
			i++ //essentially move two spaces so you dont double count
		} else {
			symbols = append(symbols, []byte{symbol})
		}
	}
	return
}

// func ConvertToRoman(arabic int) string { Previous implementation

// 	var result strings.Builder

// 	for arabic > 0 {
// 		switch {
// 		case arabic > 9:
// 			result.WriteString("X")
// 			arabic -= 10
// 		case arabic > 8:
// 			result.WriteString("IX")
// 			arabic -= 9
// 		case arabic > 4:
// 			result.WriteString("V")
// 			arabic -= 5
// 		case arabic > 3:
// 			result.WriteString("IV")
// 			arabic -= 4
// 		default:
// 			result.WriteString("I")
// 			arabic--
// 		}
// 	}

// 	return result.String()

// }
