package uuid

var (
	v1LastTimestamp int64
	v1LastSequence  uint16
)

func NewV1() (uuid UUID, err error) {
	timestamp := intervalsSinceEpoch()
	uuid[0] = byte(timestamp >> 24) // time_low 32 bits from 0 to 31
	uuid[1] = byte(timestamp >> 16)
	uuid[2] = byte(timestamp >> 8)
	uuid[3] = byte(timestamp >> 0)
	uuid[4] = byte(timestamp >> 40) // time_mid 16 bits from 32 to 47
	uuid[5] = byte(timestamp >> 32)
	uuid[6] = byte(timestamp >> 56) // time_high 12 bits from 52 to 63 (bits 48 to 51 are overwritten by version)
	uuid[7] = byte(timestamp >> 48)
	if v1LastTimestamp == timestamp {
		v1LastSequence++
	} else {
		seq, err := randomUint16()
		if err != nil {
			return uuid, err
		}
		v1LastSequence = seq
		v1LastTimestamp = timestamp
	}
	uuid[8] = byte(v1LastSequence >> 8) // clock_seq 14 bits from 66 to 79 (bits 64 and 65 are overwritten by variant)
	uuid[9] = byte(v1LastSequence >> 0)
	hwaddr, err := getHWAddr()
	if err != nil {
		return uuid, err
	}
	copy(uuid[10:], hwaddr) // node 48 bits from 80 to 127
	uuid.setVersion(1)
	return
}
