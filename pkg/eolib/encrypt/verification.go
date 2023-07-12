package encrypt

// ServerVerificationHash is used for verification that a game server is genuine.
func ServerVerificationHash(challenge int) int {
	challenge++
	return 110905 + (challenge%9+1)*((11092004-challenge)%((challenge%11+1)*119))*119 + challenge%2004
}
