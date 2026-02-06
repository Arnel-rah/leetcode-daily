type AuthenticationManager struct {
    timeToLive int
    tokens     map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
    return AuthenticationManager{
        timeToLive: timeToLive,
        tokens:     make(map[string]int),
    }
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
    this.tokens[tokenId] = currentTime + this.timeToLive
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
    expire, exists := this.tokens[tokenId]
    if exists && expire > currentTime {
        this.tokens[tokenId] = currentTime + this.timeToLive
    }
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
    count := 0
    for _, expire := range this.tokens {
        if expire > currentTime {
            count++
        }
    }
    return count
}