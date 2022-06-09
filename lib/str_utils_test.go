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
