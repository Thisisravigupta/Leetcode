func maxArea(height []int) int {
    Max := 0
    i := 0
    j := len(height) - 1
    for i < j {
       width := j - i
       heights := min(height[i], height[j])
       area := width * heights
       Max = max(Max, area)
       if height[i] < height[j] {
        i++
       } else {
        j--
       }
    }
    return Max
}