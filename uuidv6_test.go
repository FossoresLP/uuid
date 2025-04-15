package uuid

import (
	"reflect"
	"testing"
)

/*
RFC 9562 A.5. Example of a UUIDv6 Value
-------------------------------------------
field       bits value
-------------------------------------------
time_high   32   0x1EC9414C
time_mid    16   0x232A
ver          4   0x6
time_high   12   0xB00
var          2   0b10
clock_seq   14   0b11, 0x3C8
node        48   0x9F6BDECED846
-------------------------------------------
total       128
-------------------------------------------
final: 1EC9414C-232A-6B00-B3C8-9F6BDECED846
*/

func TestNewV6(t *testing.T) {
	tests := []struct {
		name     string
		testTime int64
		randUint uint32
		mac      []byte
		wantUUID UUID
		nextUUID UUID
	}{
		{
			"RFC9562",
			testVecTimeRFC,
			0x33C8,
			[]byte{0x9F, 0x6B, 0xDE, 0xCE, 0xD8, 0x46},
			UUID{0x1E, 0xC9, 0x41, 0x4C, 0x23, 0x2A, 0x6B, 0x00, 0xB3, 0xC8, 0x9F, 0x6B, 0xDE, 0xCE, 0xD8, 0x46},
			UUID{0x1E, 0xC9, 0x41, 0x4C, 0x23, 0x2A, 0x6B, 0x00, 0xB3, 0xC9, 0x9F, 0x6B, 0xDE, 0xCE, 0xD8, 0x46},
		},
		{
			"DifferentTime",
			testVecTimeCustom,
			0x33C8,
			[]byte{0x9F, 0x6B, 0xDE, 0xCE, 0xD8, 0x46},
			UUID{0x1E, 0xBB, 0x64, 0x98, 0x54, 0x7D, 0x62, 0x3F, 0xB3, 0xC8, 0x9F, 0x6B, 0xDE, 0xCE, 0xD8, 0x46},
			UUID{0x1E, 0xBB, 0x64, 0x98, 0x54, 0x7D, 0x62, 0x3F, 0xB3, 0xC9, 0x9F, 0x6B, 0xDE, 0xCE, 0xD8, 0x46},
		},
		{
			"DifferentSequence",
			testVecTimeRFC,
			0x0123,
			[]byte{0x9F, 0x6B, 0xDE, 0xCE, 0xD8, 0x46},
			UUID{0x1E, 0xC9, 0x41, 0x4C, 0x23, 0x2A, 0x6B, 0x00, 0x81, 0x23, 0x9F, 0x6B, 0xDE, 0xCE, 0xD8, 0x46},
			UUID{0x1E, 0xC9, 0x41, 0x4C, 0x23, 0x2A, 0x6B, 0x00, 0x81, 0x24, 0x9F, 0x6B, 0xDE, 0xCE, 0xD8, 0x46},
		},
		{
			"DifferentMAC",
			testVecTimeRFC,
			0x33C8,
			[]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB},
			UUID{0x1E, 0xC9, 0x41, 0x4C, 0x23, 0x2A, 0x6B, 0x00, 0xB3, 0xC8, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB},
			UUID{0x1E, 0xC9, 0x41, 0x4C, 0x23, 0x2A, 0x6B, 0x00, 0xB3, 0xC9, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB},
		},
		{
			"SequenceRollover",
			testVecTimeRFC,
			0x3FFF,
			[]byte{0x9F, 0x6B, 0xDE, 0xCE, 0xD8, 0x46},
			UUID{0x1E, 0xC9, 0x41, 0x4C, 0x23, 0x2A, 0x6B, 0x00, 0xBF, 0xFF, 0x9F, 0x6B, 0xDE, 0xCE, 0xD8, 0x46},
			UUID{0x1E, 0xC9, 0x41, 0x4C, 0x23, 0x2A, 0x6B, 0x00, 0x80, 0x00, 0x9F, 0x6B, 0xDE, 0xCE, 0xD8, 0x46},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testPrepare(tt.testTime, nil, tt.randUint, tt.mac)
			v6LastTimestamp.Store(0)
			v6LastSequence.Store(0)
			gotUUID := NewV6()
			if !reflect.DeepEqual(gotUUID, tt.wantUUID) {
				t.Errorf("NewV6() = %v, want %v for first ID", gotUUID, tt.wantUUID)
			}
			seqUUID := NewV6()
			if !reflect.DeepEqual(seqUUID, tt.nextUUID) {
				t.Errorf("NewV6() = %v, want %v for sequential ID", seqUUID, tt.nextUUID)
			}
		})
	}
}
