func isPalindrome(s string) bool {
	newStr := cleanStr(s)

	i := 0
	j := len(newStr) - 1

	for i < j {
		if newStr[i] != newStr[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func cleanStr(s string) string {
	newStr := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]

		if !((ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')) {
			continue
		}

        if ch >= 'A' && ch <= 'Z' {
            ch += 'a' - 'A'
        }

		newStr += string(ch)
	}
	return newStr
}