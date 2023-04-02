package uuid

import (
	"crypto/rand"
)

var (
	v7lastTimeStamp int64
	v7lastSequence  uint16
)

func NewV7() (uuid UUID, err error) {
	now := CurrentTime().UnixMilli()
	uuid[0] = byte(now >> 40) //1-6 bytes: big-endian unsigned number of Unix epoch timestamp
	uuid[1] = byte(now >> 32)
	uuid[2] = byte(now >> 24)
	uuid[3] = byte(now >> 16)
	uuid[4] = byte(now >> 8)
	uuid[5] = byte(now)
	if now == v7lastTimeStamp {
		v7lastSequence++
	} else {
		v7lastTimeStamp = now
		v7lastSequence = 0
	}
	uuid[6] = byte(v7lastSequence >> 8)
	uuid[7] = byte(v7lastSequence)
	_, err = rand.Read(uuid[8:])
	if err != nil {
		return
	}
	uuid.setVersion(7)
	return
}
