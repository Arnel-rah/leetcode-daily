func sortVowels(s string) string {
    vowels := []byte{}
    for _, c := range []byte(s) {
        if isVowel(c) {
            vowels = append(vowels, c)
        }
    }
    
    sort.Slice(vowels, func(i, j int) bool {
        return vowels[i] < vowels[j]
    })

    res := []byte(s)
    j := 0
    for i := range res {
        if isVowel(res[i]) {
            res[i] = vowels[j]
            j++
        }
    }
    
    return string(res)
}

func isVowel(c byte) bool {
    switch c {
    case 'a','e','i','o','u','A','E','I','O','U':
        return true
    }
    return false
}