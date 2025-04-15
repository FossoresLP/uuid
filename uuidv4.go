package uuid

// NewV4 returns a new UUID generated from cryptographically secure random data.
func NewV4() (uuid UUID) {
	randomSource(uuid[:])
	uuid.setVersion(4)
	return
}
