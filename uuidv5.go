package uuid

import (
	"crypto/sha1"
)

func NewV5(ns UUID, name string) (uuid UUID, err error) {
	hash := sha1.New()
	hash.Write(ns[:])
	hash.Write([]byte(name))
	copy(uuid[:], hash.Sum(nil))
	uuid.setVersion(5)
	return
}
