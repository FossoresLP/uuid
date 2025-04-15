package uuid

import (
	"sync/atomic"
)

var (
	v1LastTimestamp atomic.Int64
	v1LastSequence  atomic.Uint32
)

func NewV1() (uuid UUID) {
	timestamp := intervalsSinceEpoch()
	uuid[0] = byte(timestamp >> 24) // time_low 32 bits from 0 to 31
	uuid[1] = byte(timestamp >> 16)
	uuid[2] = byte(timestamp >> 8)
	uuid[3] = byte(timestamp >> 0)
	uuid[4] = byte(timestamp >> 40) // time_mid 16 bits from 32 to 47
	uuid[5] = byte(timestamp >> 32)
	uuid[6] = byte(timestamp >> 56) // time_high 12 bits from 52 to 63 (bits 48 to 51 are overwritten by version)
	uuid[7] = byte(timestamp >> 48)
	var seq uint32
	if timestamp == v1LastTimestamp.Swap(timestamp) {
		seq = v1LastSequence.Add(1)
	} else {
		seq = randomUint32(0x4000)
		v1LastSequence.Store(seq)
	}
	uuid[8] = byte(seq >> 8) // clock_seq 14 bits from 66 to 79 (bits 64 and 65 are overwritten by variant)
	uuid[9] = byte(seq >> 0)
	copy(uuid[10:], mac) // node 48 bits from 80 to 127
	uuid.setVersion(1)
	return
}
