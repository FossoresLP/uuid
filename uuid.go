package uuid

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	mrand "math/rand/v2"
	"net"
	"time"
)

// epochToUnix represents the 100ns intervals between 1582-10-15T00:00:00.00Z and 1970-01-01T00:00:00.00Z
const epochToUnix int64 = 122192928000000000

// currentTime is a function used to get the current time. It defaults to time.Now but may be replaced for testing purposes.
var currentTime func() time.Time = time.Now

// randomSource is a function used to generate random bytes. It defaults to crypto/rand.Rand but may be replaced for testing purposes.
var randomSource func([]byte) (int, error) = crand.Read

// randomUint32 is a function used to generate a random uint32. It defaults to math/rand/v2.Uint32N but may be replaced for testing purposes.
var randomUint32 func(uint32) uint32 = mrand.Uint32N

var netInterfaces func() ([]net.Interface, error) = net.Interfaces

var (
	mac net.HardwareAddr // MAC address - derived from RandomMAC or from a network card if UseHardwareMAC is true
)

func init() {
	mac = make(net.HardwareAddr, 6)
	randomSource(mac)
	mac[0] |= 0x03 // set local and multicast bits - spec requires only multicast to be set
}

// SetMACAddress sets the MAC address to be used for generating UUIDs.
// The MAC address must be 6 bytes long.
// If the MAC address is not set, a random MAC address will be generated.
// WARNING: This function is not thread-safe. Make sure to set the MAC address before generating any UUIDs.
func SetMACAddress(macAddr net.HardwareAddr) error {
	if len(macAddr) != 6 {
		return fmt.Errorf("invalid MAC address length: %d", len(macAddr))
	}
	copy(mac, macAddr)
	return nil
}

// UseHardwareMAC sets the MAC address to be used for generating UUIDs to the first valid hardware MAC address found on the system.
// If no valid hardware MAC address is found, an error is returned.
// WARNING: This function is not thread-safe. Make sure to set the MAC address before generating any UUIDs.
func UseHardwareMAC() error {
	ifaces, err := netInterfaces()
	if err != nil {
		return err
	}
	for _, iface := range ifaces {
		if len(iface.HardwareAddr) == 6 {
			mac = iface.HardwareAddr
			return nil
		}
	}
	return fmt.Errorf("no valid hardware MAC address found")
}

// UUID represents a Universal Unique Identifier as an array containing 16 bytes
type UUID [16]byte

func (uuid *UUID) setVersion(v byte) {
	uuid[6] = (uuid[6] & 0x0f) | (v << 4) // Version
	uuid[8] = (uuid[8] & 0x3f) | 0x80     // Variant 10
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
	return epochToUnix + currentTime().UTC().UnixNano()/100
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
