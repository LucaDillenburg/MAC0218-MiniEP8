package core

func IsPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		begin := i
		last := len(str) - i - 1
		if str[begin] != str[last] {
			return false
		}
	}
	return true
}
