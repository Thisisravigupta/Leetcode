func climbStairs(n int) int {
    if n == 1 || n == 2 {
        return n
    }
    a, b := 1, 2
    
    res := a
    for i := 3; i <= n; i++ {
        res = a + b
        a = b
        b = res
    }
    return res

}