package tripofgo

func ReverseString(s string) (result string) {
	// iterate from last element to first element
	for i := len(s); i > 0; i-- {
		result += s[i-1 : i]
	}
	return
}

func CountWords(s string) int {
	count := 1

	// if s is empty
	if len(s) == 0 {
		return 0
	}

	for _, v := range s {
		// 32 is ASCII for space
		if v == 32 {
			count++
		}
	}
	return count
}

func Palindrome(s string) string {
	var result string
	// reverse string
	for i := len(s); i > 0; i-- {
		result += s[i-1 : i]
	}

	// compare with early string
	if result == s {
		return s
	}

	return "Not Palindrome"
}

func ChangeWithOtherChar(from, to, s string) (result string) {
	for _, v := range s {
		// compare with string
		if from == string(v) {
			// change with other character
			result += to
		} else {
			result += string(v)
		}
	}
	return
}

func ChangeWithOthers(from, to, s string) string {
	c := 1
	var newStr string

	// if to is less than from
	if len(to) < len(from) {
		return "must be longer than early"
	}

	// iterate s
	for _, v := range s {
		// if c as index is larger than length of from string
		if c > len(from) || from[c-1:c] != string(v) {
			newStr += string(v)
		} else if from[c-1:c] == string(v) {
			// compare from string with s string
			newStr += to[c-1 : c]
			c++
		}
	}

	return newStr
}
