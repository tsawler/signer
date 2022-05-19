package signer

import (
	"fmt"
	goalone "github.com/bwmarrin/go-alone"
	"strings"
	"time"
)

// Signature is the type for the package. Secret is the signer secret, a lengthy
// and hard to guess string we use to sign things.
type Signature struct {
	Secret string
}

// GenerateTokenFromString generates a signed token and returns it
func (s *Signature) GenerateTokenFromString(data string) string {
	var urlToSign string

	str := goalone.New([]byte(s.Secret), goalone.Timestamp)
	if strings.Contains(data, "?") {
		urlToSign = fmt.Sprintf("%s&hash=", data)
	} else {
		urlToSign = fmt.Sprintf("%s?hash=", data)
	}

	tokenBytes := str.Sign([]byte(urlToSign))
	token := string(tokenBytes)

	return token
}

// VerifyToken verifies a signed token and returns true if it is valid,
// false if it is not.
func (s *Signature) VerifyToken(token string) bool {
	str := goalone.New([]byte(s.Secret), goalone.Timestamp)
	_, err := str.Unsign([]byte(token))

	if err != nil {
		// signature is not valid. Token was tampered with, forged, or maybe it's
		// not even a token at all! Either way, it's not safe to use it.
		return false
	}
	// valid hash
	return true

}

// Expired checks to see if a token has expired. It returns true if
// the token was created within minutesUntilExpire, and false otherwise.
func (s *Signature) Expired(token string, minutesUntilExpire int) bool {
	str := goalone.New([]byte(s.Secret), goalone.Timestamp)
	ts := str.Parse([]byte(token))

	// time.Duration(seconds)*time.Second
	return time.Since(ts.Timestamp) > time.Duration(minutesUntilExpire)*time.Minute
}
