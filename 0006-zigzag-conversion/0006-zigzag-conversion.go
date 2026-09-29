func convert(s string, numRows int) string {
    if numRows == 1 || numRows >= len(s) {
        return s
    }
    
    rows := make([]string, numRows)

    currentRow := 0
    goingDown := true

    for j := 0; j < len(s); j++ {

        // 1. Put s[j] into currentRow
        rows[currentRow] += string(s[j])

        // 2. If we reached bottom, what should happen?
        if currentRow == numRows - 1 {
            goingDown = false
        }

        // 3. If we reached top, what should happen?
        if currentRow == 0 {
            goingDown = true
        }

        // 4. Move currentRow
        if goingDown {
            currentRow++ 
        } else {
            currentRow--
        }
    }

    fmt.Println(rows)

    res := ""

    for i := 0; i < numRows; i++ {
        res += rows[i]
    }

    return res
}