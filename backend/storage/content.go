package storage

import (
	"crypto/sha256"
	"encoding/hex"
)

type (
	// Content represents the actual contents of a blob
	Content []byte

	// ContentID holds the identify of a content, including it's hash alg
	// and the actual value.
	//
	// ContentID should be considered immutable
	ContentID struct {
		Alg string
		ID  string
	}
)

// GetID runs the hash algorithm using content as input and then generates a new
// ContentID
func (c Content) GetID() ContentID {
	h := sha256.New()
	h.Write(c)
	sum := h.Sum(nil)
	return ContentID{
		Alg: "sha256",
		ID:  hex.EncodeToString(sum),
	}
}
