func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    original := x
    ans := 0
    for x > 0 {
        rem := x % 10
        ans = ans * 10 + rem
        x /= 10
    }
    return original == ans
}