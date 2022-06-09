package core

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		// not palindrome
		{"palindrome", false},
		{"12345321", false},
		// palindrome
		{"palinnilap", true},
		{"palinilap", true},
		{"123x321", true},
	}

	for i, test := range tests {
		got := IsPalindrome(test.param)
		if test.expected != got {
			t.Errorf("%d. Test error (expected: %t, got: %t)", i, test.expected, got)
		}
	}
}

func TestReverseString(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"1234", "4321"},
		{"hello world", "dlrow olleh"},
		{"123x321", "123x321"},
	}

	for i, test := range tests {
		got := ReverseString(test.param)
		if test.expected != got {
			t.Errorf("%d. Test error (expected: %s, got: %s)", i, test.expected, got)
		}
	}
}

func TestCountIn(t *testing.T) {
	var tests = []struct {
		paramStr  string
		paramChar rune
		expected  int
	}{
		{"1234", '5', 0},
		{"hello world", 'l', 3},
		{"this is great", ' ', 2},
	}

	for i, test := range tests {
		got := CountIn(test.paramStr, test.paramChar)
		if test.expected != got {
			t.Errorf("%d. Test error (expected: %d, got: %d)", i, test.expected, got)
		}
	}
}
