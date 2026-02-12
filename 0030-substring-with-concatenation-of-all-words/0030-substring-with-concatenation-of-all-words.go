func findSubstring(s string, words []string) []int {
    if len(words) == 0 {
        return nil
    }
    wl := len(words[0])
    total := wl * len(words)
    if len(s) < total {
        return nil
    }

    need := make(map[string]int)
    for _, w := range words {
        need[w]++
    }

    ans := []int{}

    for i := 0; i <= len(s)-total; i++ {
        seen := make(map[string]int)
        ok := true
        for j := 0; j < len(words); j++ {
            sub := s[i+j*wl : i+(j+1)*wl]
            if cnt, exists := need[sub]; !exists {
                ok = false
                break
            } else {
                seen[sub]++
                if seen[sub] > cnt {
                    ok = false
                    break
                }
            }
        }
        if ok {
            ans = append(ans, i)
        }
    }

    return ans
}