package signer

import (
	"testing"
)

func TestSignature_GenerateTokenFromString(t *testing.T) {
	s := Signature{Secret: "abc123"}

	signed := s.GenerateTokenFromString("http://example.com/test?id=1")

	if len(signed) == 0 {
		t.Error("signing failed")
	}

	signed = s.GenerateTokenFromString("http://example.com/test")

	if len(signed) == 0 {
		t.Error("signing failed")
	}
}

func TestSignature_VerifyToken(t *testing.T) {
	s := Signature{Secret: "abc123"}

	signed := s.GenerateTokenFromString("http://example.com/test?id=1")

	valid := s.VerifyToken(signed)

	if !valid {
		t.Error("valid token shows as invalid")
	}

	valid = s.VerifyToken("bad signature")

	if valid {
		t.Error("invalid token shows as valid")
	}
}

func TestSignature_Expired(t *testing.T) {
	s := Signature{Secret: "abc123"}
	signed := s.GenerateTokenFromString("http://example.com/test?id=1")
	expired := s.Expired(signed, 1)

	if expired {
		t.Error("token shows expired when it should not")
	}
}
