package uuid

import (
	"crypto/md5"
)

func NewV3(ns UUID, name string) (uuid UUID) {
	hash := md5.New()
	hash.Write(ns[:])
	hash.Write([]byte(name))
	copy(uuid[:], hash.Sum(nil))
	uuid.setVersion(3)
	return
}
