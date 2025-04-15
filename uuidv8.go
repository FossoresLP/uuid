package uuid

// NewV8 returns a new UUID based on the provided data.
// The data must be 16 bytes long.
// Bits 48-51 of the UUID are set to 0b1000 (version 8).
// Bits 64-65 are set to 0b10 (variant RFC 4122/RFC 9562).
func NewV8(data []byte) (uuid UUID) {
	copy(uuid[:], data)
	uuid.setVersion(8)
	return
}
