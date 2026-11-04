func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    chArr := make([]int, 26)

    for _, ch := range s {
        chArr[ch - 'a']++
    }

    for _, ch := range t {
        chArr[ch - 'a']--
    }

    for i := 0; i < len(chArr); i++ {
        if chArr[i] != 0 {
            return false
        }
    }
    return true
}