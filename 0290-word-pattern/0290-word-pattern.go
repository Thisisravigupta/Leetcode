func wordPattern(pattern string, s string) bool {
    words := strings.Fields(s)

    if len(pattern) != len(words) {
        return false
    }

    charToWord := make(map[byte]string)
    wordToChar := make(map[string]byte)

    for i := 0; i < len(pattern); i++ {
        char := pattern[i]
        word := words[i]

        mappedChar, exists := charToWord[char]
        if exists && mappedChar != word {
            return false
        }

        mappedWord, exists := wordToChar[word]
        if exists && mappedWord != char {
            return false
        }

        charToWord[char] = word
        wordToChar[word] = char
    }
    return true
}