func spellchecker(wordlist []string, queries []string) []string {
    exact := make(map[string]struct{})
    capMap := make(map[string]string)
    vowelMap := make(map[string]string)

    for _, w := range wordlist {
        exact[w] = struct{}{}
        low := strings.ToLower(w)
        if _, ok := capMap[low]; !ok {
            capMap[low] = w
        }
        pat := strings.Map(func(r rune) rune {
            if strings.ContainsRune("aeiou", r) {
                return '*'
            }
            return r
        }, low)
        if _, ok := vowelMap[pat]; !ok {
            vowelMap[pat] = w
        }
    }

    res := make([]string, len(queries))
    for i, q := range queries {
        if _, ok := exact[q]; ok {
            res[i] = q
            continue
        }
        low := strings.ToLower(q)
        if orig, ok := capMap[low]; ok {
            res[i] = orig
            continue
        }
        pat := strings.Map(func(r rune) rune {
            if strings.ContainsRune("aeiou", r) {
                return '*'
            }
            return r
        }, low)
        if orig, ok := vowelMap[pat]; ok {
            res[i] = orig
            continue
        }
        res[i] = ""
    }
    return res
}