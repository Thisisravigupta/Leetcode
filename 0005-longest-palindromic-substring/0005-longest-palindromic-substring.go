func longestPalindrome(s string) string {
    res := ""
    for ax := 0; ax < len(s); ax++ {
        orb := 1
        size := 1
        
        for ax - orb >= 0 && ax + orb < len(s) {
            if s[ax - orb] == s[ax + orb] {
                orb++
                size += 2
            } else {
                break
            }
        }

        if size > len(res) {
            sIdx := ax - size/2
            res = s[sIdx:sIdx+size] 
        }
    }

    for ax := 0; ax < len(s); ax++ {
        orb := 1
        size := 0
        
        for ax - orb + 1 >= 0 && ax + orb < len(s) {
            if s[ax - orb + 1] == s[ax + orb] {
                orb++
                size += 2
            } else {
                break
            }
        }

        if size > len(res) {
            sIdx := ax - size/2 + 1
            res = s[sIdx:sIdx+size] 
        }
    }

    return res
}