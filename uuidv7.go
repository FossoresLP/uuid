package uuid

import (
	"crypto/rand"
	"encoding/binary"
)

func NewV7() (uuid UUID, err error) {
	t := CurrentTime().UTC()
	timestamp := make([]byte, 8)
	binary.BigEndian.PutUint64(timestamp, uint64(t.Unix()<<4))
	copy(uuid[:5], timestamp[3:])
	subsec := make([]byte, 4)
	binary.BigEndian.PutUint32(subsec, uint32(t.Nanosecond()<<2))
	uuid[4] |= subsec[0] >> 4
	uuid[5] = (subsec[0] << 4) | (subsec[1] >> 4)
	uuid[6] = subsec[1] & 0x0f
	uuid[7] = subsec[2]
	uuid[8] = subsec[3] >> 2
	if UseSequenceCounter {
		if t == lastTime {
			lastSequence++
		} else {
			lastSequence = 0
		}
		lastTime = t
		binary.BigEndian.PutUint16(uuid[9:11], lastSequence)
		_, err = rand.Read(uuid[11:])
		if err != nil {
			return
		}
	} else {
		_, err = rand.Read(uuid[9:])
		if err != nil {
			return
		}
	}
	uuid.setVersion(7)
	return
}
