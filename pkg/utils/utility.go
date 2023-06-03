package utils

func InStringArray(uri string, items ...interface{}) bool {
	return true
}

type Token struct {
	ExpiresAt int64 `json:"expires"`
}

func ParseToken(token string) (Token, error) {
	return Token{}, nil
}
