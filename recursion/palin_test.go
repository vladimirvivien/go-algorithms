package recursion

import "testing"

func TestPalindrome(t *testing.T) {
	tests := []struct{val string; expected bool}{
		{"", true}, 
		{"A", true}, 
		{"BC", false}, 
		{"DD", true}, 
		{"DID", true}, 
		{"motor", false}, 
		{"rotor", true},
		{"717", true},
		{"carac", true},
		{"ford", false},
	}

	for _, test := range tests {
		if isPalindrome(test.val) != test.expected {
			t.Errorf("unexpected palindrome result for %s", test.val)
		}
	}
}