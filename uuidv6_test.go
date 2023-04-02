package uuid

import (
	"testing"
)

func TestNewV6(t *testing.T) {
	CurrentTime = testTime()
	testRand(0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF)
	id, err := NewV6()
	t.Log(id)
	if err == nil || err.Error() != "version 6 not supported, yet" {
		t.Error("UUIDv6 did not return an error - please implement tests")
	}
}
