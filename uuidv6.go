package uuid

import (
	"crypto/rand"
	"encoding/binary"
)

func NewV6() (uuid UUID, err error) {

	t := make([]byte, 8)
	diffToEpoch := uint64(122192928000000000 + CurrentTime().UTC().UnixNano()/100)
	binary.BigEndian.PutUint64(t, diffToEpoch<<4)
	copy(uuid[:6], t[:6])
	uuid[6] = t[6] >> 4
	uuid[7] = (t[6] << 4) | (t[7] >> 4)
	if diffToEpoch == lastDiffToEpoch {
		sequenceCounterEpoch++
	} else {
		sequenceCounterEpoch = 0
	}
	binary.BigEndian.PutUint16(uuid[8:10], sequenceCounterEpoch)
	_, err = rand.Read(uuid[10:])
	if err != nil {
		return
	}
	uuid.setVersion(6)
	return
}