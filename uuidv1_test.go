package uuid

import (
	"testing"
)

func TestNewV1(t *testing.T) {
	_, err := NewV1()
	if err == nil || err.Error() != "version 1 not supported, yet" {
		t.Error("UUIDv1 did not return an error - please implement tests")
	}
}
