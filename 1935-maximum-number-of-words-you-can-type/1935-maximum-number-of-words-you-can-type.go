func canBeTypedWords(text string, brokenLetters string) int {
    broken := [26]bool{}
    for _, c := range brokenLetters {
        broken[c-'a'] = true
    }

    ans := 0
    words := strings.Split(text, " ")

    for _, word := range words {
        can := true
        for _, c := range word {
            if broken[c-'a'] {
                can = false
                break
            }
        }
        if can {
            ans++
        }
    }

    return ans
}