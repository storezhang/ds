package stringx

// Palindrome 回文字符串判断
func Palindrome(str string) (palindrome bool) {
	left := 0
	right := len(str) - 1
	for left < right {
		if str[left] != str[right] {
			return
		} else {
			left++
			right--
		}
	}
	palindrome = true

	return
}
