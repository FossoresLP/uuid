package uuid

var (
	v6LastTimestamp int64
	v6LastSequence  uint16
)

func NewV6() (uuid UUID, err error) {
	timestamp := intervalsSinceEpoch()
	uuid[0] = byte(timestamp >> 52) // time_high 32 bits from 0 to 31
	uuid[1] = byte(timestamp >> 44)
	uuid[2] = byte(timestamp >> 36)
	uuid[3] = byte(timestamp >> 28)
	uuid[4] = byte(timestamp >> 20) // time_mid 16 bits from 32 to 47
	uuid[5] = byte(timestamp >> 12)
	uuid[6] = byte(timestamp >> 8) // time_low 12 bits from 52 to 63 (bits 48 to 51 are overwritten by version)
	uuid[7] = byte(timestamp >> 0)
	if v6LastTimestamp == timestamp {
		v6LastSequence++
	} else {
		seq, err := randomUint16()
		if err != nil {
			return uuid, err
		}
		v6LastSequence = seq
		v6LastTimestamp = timestamp
	}
	uuid[8] = byte(v6LastSequence >> 8) // clock_seq 14 bits from 66 to 79 (bits 64 and 65 are overwritten by variant)
	uuid[9] = byte(v6LastSequence >> 0)
	hwaddr, err := getHWAddr()
	if err != nil {
		return uuid, err
	}
	copy(uuid[10:], hwaddr) // node 48 bits from 80 to 127
	uuid.setVersion(6)
	return
}
