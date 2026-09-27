func firstUniqChar(s string) int {
    letters := make([]int, 26)

    for i := 0; i < len(s); i++ {
        ch := s[i] - 'a'
        letters[ch]++
    }

    for i := 0; i < len(s); i++ {
        ch := s[i] - 'a'
        if letters[ch] == 1 {
            return i
        }
    }

    fmt.Println(letters)

    return -1
}