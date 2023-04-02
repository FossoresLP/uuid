package uuid

import (
	"encoding/binary"
	"time"
)

// CurrentTime is a function used to get the current time. It defaults to time.Now but can be set to a different function in case a different time source should be used.
var CurrentTime func() time.Time = time.Now

// UUID represents a Universal Unique Identifier as an array containing 16 bytes
type UUID [16]byte

func (uuid *UUID) setVersion(v byte) {
	uuid[6] = (uuid[6] & 0x0f) | (v << 4) // Version
	uuid[8] = (uuid[8] & 0x3f) | 0x80     // Variant 10
}

// Must wraps the output of New and panics when an error occured
func Must(uuid UUID, err error) UUID {
	if err != nil {
		panic(err)
	}
	return uuid
}

// MustString works like Must but returns a string immediately
func MustString(uuid UUID, err error) string {
	if err != nil {
		panic(err)
	}
	return uuid.String()
}

// IsNil returns true if the UUID contains only zeros and is therefore empty and invalid
func (uuid UUID) IsNil() bool {
	return uuid == [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
}

// IsMax returns true if the UUID contains only ones and is therefore invalid
func (uuid UUID) IsMax() bool {
	return uuid == [16]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
}

// Version returns the version of the UUID
func (uuid UUID) Version() int {
	return int(uuid[6] >> 4)
}

// Timestamp returns the timestamp of the UUID or nil if the UUID does not contain a timestamp
func (uuid UUID) Timestamp() time.Time {
	switch uuid.Version() {
	case 7:
		i := int64(binary.BigEndian.Uint64(append([]byte{0x00, 0x00}, uuid[:6]...)))
		return time.Unix(i/1000, (i%1000)*1000000)
	default:
		return time.Unix(0, 0)
	}
}
