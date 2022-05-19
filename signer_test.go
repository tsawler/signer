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

func TestSignature_VerifyToken(t *testing.T) {
	sign := Signature{Secret: "abc123"}

	signed, _ := sign.SignURL("https://example.com/test?id=1")

	valid, err := sign.VerifyURL(signed)
	if err != nil {
		t.Error("error when validating url")
	}
	if !valid {
		t.Error("valid token shows as invalid")
	}

	valid, err = sign.VerifyURL("https://www.example.com/some/url")
	if valid {
		t.Error("invalid token shows as valid")
	}

	valid, err = sign.VerifyURL("not a url")
	if err == nil {
		t.Error("no error when validating non-url")
	}
	if valid {
		t.Error("returned valid on non url")
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
