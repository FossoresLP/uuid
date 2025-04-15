package uuid

func NewV4() (uuid UUID) {
	randomSource(uuid[:])
	uuid.setVersion(4)
	return
}
