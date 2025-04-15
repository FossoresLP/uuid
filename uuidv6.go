package uuid

import (
	"sync/atomic"
)

var (
	v6LastTimestamp atomic.Int64
	v6LastSequence  atomic.Uint32
)

// NewV6 returns a new UUID based on the current timestamp and MAC address.
// The timestamp is retrieved from the system clock.
// The MAC address is randomly generated at application startup and can be overridden using SetMACAddress or UseHardwareMAC.
// Unlike UUIDv1, UUIDv6 is designed to be sortable by time using binary or lexicographical comparison.
func NewV6() (uuid UUID) {
	timestamp := intervalsSinceEpoch()
	uuid[0] = byte(timestamp >> 52) // time_high 32 bits from 0 to 31
	uuid[1] = byte(timestamp >> 44)
	uuid[2] = byte(timestamp >> 36)
	uuid[3] = byte(timestamp >> 28)
	uuid[4] = byte(timestamp >> 20) // time_mid 16 bits from 32 to 47
	uuid[5] = byte(timestamp >> 12)
	uuid[6] = byte(timestamp >> 8) // time_low 12 bits from 52 to 63 (bits 48 to 51 are overwritten by version)
	uuid[7] = byte(timestamp >> 0)
	var seq uint32
	if timestamp == v6LastTimestamp.Swap(timestamp) {
		seq = v6LastSequence.Add(1)
	} else {
		seq = randomUint32(0x4000)
		v6LastSequence.Store(seq)
	}
	uuid[8] = byte(seq >> 8) // clock_seq 14 bits from 66 to 79 (bits 64 and 65 are overwritten by variant)
	uuid[9] = byte(seq >> 0)
	copy(uuid[10:], mac) // node 48 bits from 80 to 127
	uuid.setVersion(6)
	return
}
