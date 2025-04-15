package uuid

// NewV7 returns a new UUID based on the current timestamp and random data.
// The timestamp is retrieved from the system clock.
// The random data is generated using the cryptographically secure random number generator.
// This implementation uses the fractional millisecond approach for ordering of UUIDs within the same millisecond.
func NewV7() (uuid UUID) {
	time := currentTime()
	ms := time.UnixMilli()
	uuid[0] = byte(ms >> 40) //1-6 bytes: 48-bit big-endian unsigned number of Unix epoch timestamp
	uuid[1] = byte(ms >> 32)
	uuid[2] = byte(ms >> 24)
	uuid[3] = byte(ms >> 16)
	uuid[4] = byte(ms >> 8)
	uuid[5] = byte(ms)

	frac := uint16(time.Nanosecond() % 1000000 * 4095 / 999999)

	uuid[6] = byte(frac >> 8) //7-8 bytes: 12-bit big-endian fractional part of Unix epoch timestamp
	uuid[7] = byte(frac)
	randomSource(uuid[8:]) // 9-16 bytes: 64-bit cryptographically random data
	uuid.setVersion(7)
	return
}
