package tokens

import (
	"time"
	"crypto/sha256"
	"crypto/rand"
	"encoding/base32"
)

const (
	ScopeAuth = "authentication"
)

type Token struct {
	Plaintext string `json:"token"`
	Hash []byte `json:"-"`
	UserID int `json:"-"`
	Expire time.Time `json:"expire"`
	Scope string `json:"-"`
}

func GenerateToken(userID int, ttl time.Duration, scope string) (*Token, error) {
	token := &Token{
		UserID: userID,
		Expire: time.Now().Add(ttl),
		Scope: scope,
	}

	emptyBytes := make([]byte, 32)
	_, err := rand.Read(emptyBytes)
	if err != nil {
		return nil, err
	}

	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(emptyBytes)
	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]
	return token, nil
}
