package uuid

import (
	"testing"
)

func TestNewV8(t *testing.T) {
	_, err := NewV8()
	if err == nil || err.Error() != "version 8 not supported, yet" {
		t.Error("UUIDv8 did not return an error - please implement tests")
	}
}
