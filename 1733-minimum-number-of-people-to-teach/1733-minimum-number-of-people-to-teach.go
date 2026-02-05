func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
    m := len(languages)
    knows := make([]map[int]bool, m+1)
    
    for i := 1; i <= m; i++ {
        knows[i] = make(map[int]bool)
        for _, l := range languages[i-1] {
            knows[i][l] = true
        }
    }

    problematic := [][2]int{}
    for _, f := range friendships {
        u, v := f[0], f[1]
        common := false
        for lang := range knows[u] {
            if knows[v][lang] {
                common = true
                break
            }
        }
        if !common {
            problematic = append(problematic, [2]int{u, v})
        }
    }

    if len(problematic) == 0 {
        return 0
    }

    ans := m + 1
    for lang := 1; lang <= n; lang++ {
        need := make(map[int]bool)
        for _, p := range problematic {
            if !knows[p[0]][lang] {
                need[p[0]] = true
            }
            if !knows[p[1]][lang] {
                need[p[1]] = true
            }
        }
        if len(need) < ans {
            ans = len(need)
        }
    }
    return ans
}