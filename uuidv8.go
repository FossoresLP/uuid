package uuid

func NewV8(data []byte) (uuid UUID) {
	copy(uuid[:], data)
	uuid.setVersion(8)
	return
}
