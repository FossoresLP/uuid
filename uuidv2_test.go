package uuid

import (
	"testing"
)

func TestNewV2(t *testing.T) {
	_, err := NewV2()
	if err == nil {
		t.Error("uuid.NewV2() did not generate an error even though it is not supported")
	}
}
