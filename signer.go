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

// SignURL generates a signed url and returns it, stripping off http:// and https://
func (s *Signature) SignURL(data string) string {
	var urlToSign string

	exploded := strings.Split(data, "//")

	pen := goalone.New([]byte(s.Secret), goalone.Timestamp)

	if strings.Contains(exploded[1], "?") {
		// handle case where URL contains query parameters
		urlToSign = fmt.Sprintf("%s&hash=", exploded[1])
	} else {
		// no query parameters
		urlToSign = fmt.Sprintf("%s?hash=", exploded[1])
	}

	tokenBytes := pen.Sign([]byte(urlToSign))
	token := string(tokenBytes)

	return fmt.Sprintf("%s//%s", exploded[0], token)
}

// VerifyURL verifies a signed url and returns true if it is valid,
// false if it is not. Note that http:// and https:// are stripped off
// before verification
func (s *Signature) VerifyURL(data string) bool {
	exploded := strings.Split(data, "//")
	pen := goalone.New([]byte(s.Secret), goalone.Timestamp)

	_, err := pen.Unsign([]byte(exploded[1]))

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
func (s *Signature) Expired(data string, minutesUntilExpire int) bool {
	exploded := strings.Split(data, "//")

	pen := goalone.New([]byte(s.Secret), goalone.Timestamp)
	ts := pen.Parse([]byte(exploded[1]))

	return time.Since(ts.Timestamp) > time.Duration(minutesUntilExpire)*time.Minute
}
