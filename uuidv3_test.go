package uuid

import (
	"testing"
)

/*
RFC 9562 A.2. Example of a UUIDv3 Value

Namespace (DNS):  6ba7b810-9dad-11d1-80b4-00c04fd430c8
Name:             www.example.com
------------------------------------------------------
MD5:              5df418813aed051548a72f4a814cf09e
Figure 17: UUIDv3 Example MD5
-------------------------------------------
field     bits value
-------------------------------------------
md5_high  48   0x5df418813aed
ver        4   0x3
md5_mid   12   0x515
var        2   0b10
md5_low   62   0b00, 0x8a72f4a814cf09e
-------------------------------------------
total     128
-------------------------------------------
final: 5df41881-3aed-3515-88a7-2f4a814cf09e
*/

func TestNewV3(t *testing.T) {
	want := UUID{0x5D, 0xF4, 0x18, 0x81, 0x3A, 0xED, 0x35, 0x15, 0x88, 0xA7, 0x2F, 0x4A, 0x81, 0x4C, 0xF0, 0x9E}
	id := NewV3(NamespaceDNS(), "www.example.com")
	if id != want {
		t.Errorf("uuid.NewV3() = %v, want %v", id, want)
	}
}
