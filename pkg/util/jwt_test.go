package util

import "testing"

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("admin", "123456")
	if err != nil {
		t.Fatalf("Generate token error: %v\n", err)
	}
	t.Fatalf("Generate token: %v\n", token)
}
