package signer

import (
	"strings"
	"testing"
)

var signTests = []struct {
	name     string
	url      string
	validURL bool
	hasError bool
}{
	{
		name:     "signable query params",
		url:      "https://example.com/test?id=1",
		validURL: true,
		hasError: false,
	},
	{
		name:     "signable long path query params",
		url:      "https://example.com/password/reset/finish?email=my@email.com",
		validURL: true,
		hasError: false,
	},
	{
		name:     "signable no query params",
		url:      "https://example.com/test",
		validURL: true,
		hasError: false,
	},
	{
		name:     "empty url",
		url:      "",
		validURL: false,
		hasError: true,
	},
	{
		name:     "not url",
		url:      "fish",
		validURL: false,
		hasError: true,
	},
}

func TestSignature_SignURL(t *testing.T) {
	sign := Signature{Secret: "abc123"}

	for _, e := range signTests {
		signed, err := sign.SignURL(e.url)

		if e.validURL && !strings.Contains(signed, e.url) {
			t.Errorf("%s: was not returned correctly", e.url)
		}
		if err == nil && e.hasError {
			t.Errorf("%s: does not have error, and should", e.name)
		}

		if err != nil && !e.hasError {
			t.Errorf("%s: has error, and should not have one", e.name)
		}

		if len(signed) > 0 && len(e.url) != 0 && e.validURL && e.hasError {
			t.Errorf("%s: failed to sign non-empty, valid url", e.name)
		}

		if !e.validURL && err == nil {
			t.Errorf("%s: signed non valid url", e.name)
		}
	}
}

var verifyTests = []struct {
	name       string
	url        string
	validURL   bool
	shouldPass bool
}{
	{
		name:       "valid url and sig",
		url:        "https://example.com/test?id=1",
		shouldPass: true,
		validURL:   true,
	},
	{
		name:       "valid url and invalid sig",
		url:        "https://www.example.com/some/url",
		shouldPass: false,
		validURL:   false,
	},
	{
		name:       "not a url",
		url:        "not a url",
		shouldPass: false,
		validURL:   false,
	},
}

func TestSignature_VerifyToken(t *testing.T) {
	sign := Signature{Secret: "abc123"}

	for _, e := range verifyTests {
		var signed string

		if e.validURL {
			signed, _ = sign.SignURL(e.url)
		} else {
			signed = e.url
		}

		valid, err := sign.VerifyURL(signed)

		if err != nil && e.validURL {
			t.Errorf("%s: error when validating url %s", e.name, e.url)
		}
		if !valid && e.shouldPass {
			t.Errorf("%s: valid token shows as invalid", e.name)
		}
		if valid && !e.validURL {
			t.Errorf("%s: returned valid on non url %s", e.name, e.url)
		}
	}
}

func TestSignature_Expired(t *testing.T) {
	sign := Signature{Secret: "abc123"}

	signed, _ := sign.SignURL("http://example.com/test?id=1")

	expired := sign.Expired(signed, 1)

	if !expired {
		t.Error("token shows expired when it should not")
	}

	expired = sign.Expired(signed, -1)
	if expired {
		t.Error("token shows that it is not expired when it should be")
	}
}

func TestSignature_ExpiredSeconds(t *testing.T) {
	sign := Signature{Secret: "abc123"}

	signed, _ := sign.SignURL("http://example.com/test?id=1")

	expired := sign.ExpiredSeconds(signed, 1)

	if !expired {
		t.Error("token shows expired when it should not")
	}

	expired = sign.ExpiredSeconds(signed, -1)
	if expired {
		t.Error("token shows that it is not expired when it should be")
	}
}
