package signer

import (
	"fmt"
	danger "github.com/tsawler/itsdangerous"
	"net/url"
	"strings"
	"time"
)

// Signature is the type for the package. Secret is the signer secret, a lengthy
// and hard to guess string we use to sign things. The secret must not exceed 64 characters.
type Signature struct {
	Secret string
}

// SignURL generates a signed url and returns it, stripping off http:// and https://.
func (s *Signature) SignURL(data string) (string, error) {
	// Verify that the string is a url.
	u, err := url.ParseRequestURI(data)
	if err != nil {
		return "", err
	}

	var urlToSign string
	q := ""
	if u.RawQuery != "" {
		q = "?"
	}
	stringToSign := fmt.Sprintf("%s%s%s", u.Path, q, u.RawQuery)

	pen := danger.New([]byte(s.Secret), danger.Timestamp)

	if strings.Contains(stringToSign, "?") {
		// handle case where URL contains query parameters
		urlToSign = fmt.Sprintf("%s&hash=", stringToSign)
	} else {
		// no query parameters
		urlToSign = fmt.Sprintf("%s?hash=", stringToSign)
	}

	tokenBytes := pen.Sign([]byte(urlToSign))
	token := string(tokenBytes)

	return fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, token), nil
}

// VerifyURL verifies a signed url and returns true if it is valid,
// false if it is not. Note that http:// and https:// are stripped off
// before verification.
func (s *Signature) VerifyURL(data string) (bool, error) {
	u, err := url.ParseRequestURI(data)
	if err != nil {
		return false, err
	}
	q := ""
	if u.RawQuery != "" {
		q = "?"
	}
	stringToVerify := fmt.Sprintf("%s%s%s", u.Path, q, u.RawQuery)

	pen := danger.New([]byte(s.Secret), danger.Timestamp)

	_, err = pen.Unsign([]byte(stringToVerify))
	if err != nil {
		// signature is not valid. Token was tampered with, forged, or maybe it's
		// not even a token at all! Either way, it's not safe to use it.
		return false, err
	}

	// valid hash
	return true, nil

}

// Expired checks to see if a token has expired. It returns false if
// the token was created within minutesUntilExpire, and true otherwise.
func (s *Signature) Expired(data string, minutesUntilExpire int) bool {
	exploded := strings.Split(data, "//")

	pen := danger.New([]byte(s.Secret), danger.Timestamp)
	ts := pen.Parse([]byte(exploded[1]))

	return time.Since(ts.Timestamp) < time.Duration(minutesUntilExpire)*time.Minute
}

// ExpiredSeconds checks to see if a token has expired. It returns false if
// the token was created within secondsUntilExpire, and true otherwise.
func (s *Signature) ExpiredSeconds(data string, secondsUntilExpire int) bool {
	exploded := strings.Split(data, "//")

	pen := danger.New([]byte(s.Secret), danger.Timestamp)
	ts := pen.Parse([]byte(exploded[1]))

	return time.Since(ts.Timestamp) < time.Duration(secondsUntilExpire)*time.Second
}
