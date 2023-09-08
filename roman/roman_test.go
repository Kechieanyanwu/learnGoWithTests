package roman

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q want %q", got, test.Roman)
			}
		})

	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%s gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	t.Run("asserting functions return the same value", func(t *testing.T) {
		assertion := func(arabic uint16) bool {
			if arabic > 3999 {
				// log.Println(arabic)
				return true
			}
			t.Log("testing", arabic) //helps give visibility on what we are testing
			roman := ConvertToRoman(arabic)
			fromRoman := ConvertToArabic(roman)
			return fromRoman == arabic
		}

		if err := quick.Check(assertion, nil); err != nil {
			t.Error("failed checks", err)
		}
	})

	t.Run("asserting no more than 3 consecutive symbols", func(t *testing.T) {
		var roman string
		assertion := func(arabic uint16) bool {
			if arabic > 3999 {
				// log.Println(arabic)
				return true
			}
			t.Log("testing", arabic)
			roman = ConvertToRoman(arabic)
			return checkNoMoreThanThree(t, roman)
		}

		if err := quick.Check(assertion, nil); err != nil {
			t.Error("failed checks", roman, err)
		}
	})

	t.Run("asserting only correct subtractors used", func(t *testing.T) {
		assertion := func(arabic uint16) bool {
			if arabic > 3999 {
				// log.Println(arabic)
				return true
			}
			t.Log("testing", arabic)
			roman := ConvertToRoman(arabic)
			return checkValidSubtractors(t, roman)
		}

		if err := quick.Check(assertion, nil); err != nil {
			t.Error("failed checks", err)
		}
	})

}

func checkNoMoreThanThree(t testing.TB, roman string) bool {
	t.Helper()
	for i := 0; i < (len(roman) - 3); i++ {
		if roman[i] == roman[i+1] && roman[i] == roman[i+2] && roman[i] == roman[i+3] {
			return false
		}
	}
	return true
}

func checkValidSubtractors(t testing.TB, roman string) bool {
	t.Helper()
	if len(roman) < 2 {
		return true
	}
	for i := 0; i < len(roman)-1; i++ {
		nextChar := roman[i+1]
		if roman[i] == 'C' && nextChar == 'M' || nextChar == 'D' || nextChar == 'L' || nextChar == 'X' || nextChar == 'V' || nextChar == 'I' {
			return true
		}
		if roman[i] == 'X' && nextChar == 'C' || nextChar == 'L' || nextChar == 'X' || nextChar == 'V' || nextChar == 'I' {
			return true
		}
		if roman[i] == 'I' && nextChar == 'X' || nextChar == 'V' || nextChar == 'I' {
			return true
		}
	}
	return false
}

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}
