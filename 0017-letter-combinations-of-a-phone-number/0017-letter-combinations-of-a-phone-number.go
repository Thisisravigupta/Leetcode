func letterCombinations(digits string) []string {
    mapping := map[byte]string{
        '2' : "abc",
        '3' : "def",
        '4' : "ghi",
        '5' : "jkl",
        '6' : "mno",
        '7' : "pqrs",
        '8' : "tuv",
        '9' : "wxyz",
    }

    combos := []string{""}

    for i := 0; i < len(digits); i++ {
        letters := mapping[digits[i]] // abc
        next := []string{}
        for _, combo := range combos {
            for j := 0; j < len(letters); j++ {
                next = append(next, combo+string(letters[j]))
            }
        }
        combos = next
    }

    return combos
}