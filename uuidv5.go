package uuid

import (
	"crypto/sha1"
)

// NewV5 returns a new UUID based on the SHA-1 hash of the provided namespace and name.
func NewV5(ns UUID, name string) (uuid UUID) {
	hash := sha1.New()
	hash.Write(ns[:])
	hash.Write([]byte(name))
	copy(uuid[:], hash.Sum(nil))
	uuid.setVersion(5)
	return
}
