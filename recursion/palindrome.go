package recursion

func isPalindrome(s string) bool {
	if s == "" || len(s) == 1 {
		return true
	}
	if s[0] != s[len(s)-1] {
		return false
	}
	return isPalindrome(s[1:len(s)-1])
}
