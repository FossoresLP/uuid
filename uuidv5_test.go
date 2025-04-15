package uuid

import (
	"testing"
)

/*
RFC 9562 A.4. Example of a UUIDv5 Value

Namespace (DNS):  6ba7b810-9dad-11d1-80b4-00c04fd430c8
Name:             www.example.com
----------------------------------------------------------
SHA-1:            2ed6657de927468b55e12665a8aea6a22dee3e35
Figure 22: UUIDv5 Example SHA-1
-------------------------------------------
field      bits value
-------------------------------------------
sha1_high  48   0x2ed6657de927
ver         4   0x5
sha1_mid   12   0x68b
var         2   0b10
sha1_low   62   0b01, 0x5e12665a8aea6a2
-------------------------------------------
total      128
-------------------------------------------
final: 2ed6657d-e927-568b-95e1-2665a8aea6a2
*/

func TestNewV5(t *testing.T) {
	want := UUID{0x2E, 0xD6, 0x65, 0x7D, 0xE9, 0x27, 0x56, 0x8B, 0x95, 0xE1, 0x26, 0x65, 0xA8, 0xAE, 0xA6, 0xA2}
	id := NewV5(NamespaceDNS(), "www.example.com")
	if id != want {
		t.Errorf("uuid.NewV5() = %v, want %v", id, want)
	}
}
