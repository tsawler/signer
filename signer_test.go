package signer

import (
	"testing"
)

func TestSignature_SignURL(t *testing.T) {
	sign := Signature{Secret: "abc123"}

	signed, err := sign.SignURL("https://example.com/test?id=1")
	if err != nil {
		t.Error("error validating url")
	}

	if len(signed) == 0 {
		t.Error("signing failed")
	}

	signed, _ = sign.SignURL("https://example.com/test")

	if len(signed) == 0 {
		t.Error("signing failed")
	}

	_, err = sign.SignURL("not a url")
	if err == nil {
		t.Error("invalid url did not throw an error")
	}
}

var verifyTests = []struct {
	name       string
	url        string
	validUrl   bool
	shouldPass bool
}{
	{
		name:       "valid url and sig",
		url:        "https://example.com/test?id=1",
		shouldPass: true,
		validUrl:   true,
	},
	{
		name:       "valid url and invalid sig",
		url:        "https://www.example.com/some/url",
		shouldPass: false,
		validUrl:   false,
	},
	{
		name:       "not a url",
		url:        "not a url",
		shouldPass: false,
		validUrl:   false,
	},
}

func TestSignature_VerifyToken(t *testing.T) {
	sign := Signature{Secret: "abc123"}

	for _, e := range verifyTests {
		var signed = ""

		if e.validUrl {
			signed, _ = sign.SignURL(e.url)
		} else {
			signed = e.url
		}

		valid, err := sign.VerifyURL(signed)

		if err != nil && e.validUrl {
			t.Errorf("%s: error when validating url %s", e.name, e.url)
		}
		if !valid && e.shouldPass {
			t.Errorf("%s: valid token shows as invalid", e.name)
		}
		if valid && !e.validUrl {
			t.Errorf("%s: returned valid on non url %s", e.name, e.url)
		}
	}
}

func TestSignature_Expired(t *testing.T) {
	sign := Signature{Secret: "abc123"}

	signed, _ := sign.SignURL("http://example.com/test?id=1")

	expired := sign.Expired(signed, 1)

	if expired {
		t.Error("token shows expired when it should not")
	}
}
