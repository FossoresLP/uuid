package uuid

import (
	"crypto/md5"
)

// NewV3 returns a new UUID based on the MD5 hash of the provided namespace and name.
func NewV3(ns UUID, name string) (uuid UUID) {
	hash := md5.New()
	hash.Write(ns[:])
	hash.Write([]byte(name))
	copy(uuid[:], hash.Sum(nil))
	uuid.setVersion(3)
	return
}
