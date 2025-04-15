package uuid

import (
	"testing"
)

/*
RFC 9562 A.3. Example of a UUIDv4 Value
-------------------------------------------
field     bits value
-------------------------------------------
random_a  48   0x919108f752d1
ver        4   0x4
random_b  12   0x320
var        2   0b10
random_c  62   0b01, 0xbacf847db4148a8
-------------------------------------------
total     128
-------------------------------------------
final: 919108f7-52d1-4320-9bac-f847db4148a8
*/

func TestNewV4(t *testing.T) {
	testPrepare(0, []byte{0x91, 0x91, 0x08, 0xF7, 0x52, 0xD1, 0x33, 0x20, 0x5B, 0xAC, 0xF8, 0x47, 0xDB, 0x41, 0x48, 0xA8}, 0, nil)
	want := UUID{0x91, 0x91, 0x08, 0xF7, 0x52, 0xD1, 0x43, 0x20, 0x9B, 0xAC, 0xF8, 0x47, 0xDB, 0x41, 0x48, 0xA8}
	id := NewV4()

	if id != want {
		t.Errorf("uuid.NewV4() = %v, want %v", id, want)
	}
}
