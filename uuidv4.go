package uuid

import (
	"crypto/rand"
)

func NewV4() (uuid UUID, err error) {
	_, err = rand.Read(uuid[:])
	if err != nil {
		return
	}
	uuid.setVersion(4)
	return
}
