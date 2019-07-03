package storage

import (
	"testing"
)

func TestGetID(t *testing.T) {
	content := Content([]byte("hello world"))
	cid := content.GetID()
	if cid != (ContentID{Alg: "sha256", ID: "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"}) {
		t.Fatal("invalid hash")
	}
}
