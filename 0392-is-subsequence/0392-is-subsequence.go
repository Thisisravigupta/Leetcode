func isSubsequence(s string, t string) bool {
    sIdx := 0
    for i := 0; i < len(t) && sIdx < len(s); i++ {
        if t[i] == s[sIdx] {
            sIdx++
        }
    }
    return sIdx == len(s)
}