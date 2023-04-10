package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"net"
	"time"
)

// EpochToUnix represents the 100ns intervals between 1582-10-15T00:00:00.00Z and 1970-01-01T00:00:00.00Z
const EpochToUnix int64 = 122192928000000000

// CurrentTime is a function used to get the current time. It defaults to time.Now but can be set to a different function in case a different time source should be used.
var CurrentTime func() time.Time = time.Now

var (
	UseHardwareMAC bool             = false // UseHardwareMAC defines whether to use a MAC address from a network card if available or generate one. It defaults to false for privacy reasons and because the MAC address lookup creates additional latency.
	RandomMAC      net.HardwareAddr         // Random MAC address - will be automatically generated unless set manually
)

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

func intervalsSinceEpoch() int64 {
	return EpochToUnix + CurrentTime().UTC().UnixNano()/100
}

func getHWAddr() (net.HardwareAddr, error) {
	if UseHardwareMAC {
		interfaces, err := net.Interfaces()
		if err == nil {
			for _, intf := range interfaces {
				if len(intf.HardwareAddr) == 6 {
					return intf.HardwareAddr, nil
				}
			}
		}
	}
	if RandomMAC != nil {
		return RandomMAC, nil
	}
	addr := make(net.HardwareAddr, 6)
	_, err := rand.Read(addr)
	if err != nil {
		return nil, err
	}
	addr[0] |= 0x03 // set local and multicast bits - spec requires only multicast to be set
	RandomMAC = addr
	return addr, nil
}

func randomUint16() (uint16, error) {
	bytes := make([]byte, 2)
	_, err := rand.Read(bytes)
	if err != nil {
		return 0, err
	}
	return uint16(bytes[0])<<8 + uint16(bytes[1]), nil
}

func NamespaceDNS() UUID {
	return UUID{0x6B, 0xA7, 0xB8, 0x10, 0x9D, 0xAD, 0x11, 0xD1, 0x80, 0xB4, 0x00, 0xC0, 0x4F, 0xD4, 0x30, 0xC8}
}

func NamespaceURL() UUID {
	return UUID{0x6B, 0xA7, 0xB8, 0x11, 0x9D, 0xAD, 0x11, 0xD1, 0x80, 0xB4, 0x00, 0xC0, 0x4F, 0xD4, 0x30, 0xC8}
}

func NamespaceOID() UUID {
	return UUID{0x6B, 0xA7, 0xB8, 0x12, 0x9D, 0xAD, 0x11, 0xD1, 0x80, 0xB4, 0x00, 0xC0, 0x4F, 0xD4, 0x30, 0xC8}
}

func NamespaceX500() UUID {
	return UUID{0x6B, 0xA7, 0xB8, 0x14, 0x9D, 0xAD, 0x11, 0xD1, 0x80, 0xB4, 0x00, 0xC0, 0x4F, 0xD4, 0x30, 0xC8}
}
