package signer

import (
	"testing"
)

func TestSignature_SignURL(t *testing.T) {
	sign := Signature{Secret: "abc123"}

	signed := sign.SignURL("http://example.com/test?id=1")

	if len(signed) == 0 {
		t.Error("signing failed")
	}

	signed = sign.SignURL("http://example.com/test")

	if len(signed) == 0 {
		t.Error("signing failed")
	}
}

func TestSignature_VerifyToken(t *testing.T) {
	sign := Signature{Secret: "abc123"}

	signed := sign.SignURL("http://example.com/test?id=1")

	valid := sign.VerifyURL(signed)

	if !valid {
		t.Error("valid token shows as invalid")
	}

	valid = sign.VerifyURL("http://www.baddomain.com/some/url")

	if valid {
		t.Error("invalid token shows as valid")
	}
}

func TestSignature_Expired(t *testing.T) {
	sign := Signature{Secret: "abc123"}

	signed := sign.SignURL("http://example.com/test?id=1")
	expired := sign.Expired(signed, 1)

	if expired {
		t.Error("token shows expired when it should not")
	}
}
