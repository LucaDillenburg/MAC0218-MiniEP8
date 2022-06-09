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

func ReverseString(str string) string {
	reversed := make([]byte, 0)
	for i := 0; i < len(str); i++ {
		reversed = append(reversed, str[len(str)-1-i])
	}
	return string(reversed)
}
