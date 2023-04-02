package uuid

import (
	"testing"
)

func TestNewV4(t *testing.T) {
	idv4 := UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x46, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}
	testRand(idv4[:]...)
	id, err := NewV4()
	if err != nil {
		t.Error("New failed to generate a UUIDv4")
	}

	if id != idv4 {
		t.Errorf("uuid.NewV4() = %v, want %v", id, idv4)
	}
	if id.Version() != 4 {
		t.Errorf("uuid.NewV4() generated UUID with wrong version %d", id.Version())
	}
	testRand()
	_, err = NewV4()
	if err == nil {
		t.Error("New did not fail when no random data was available")
	}
}
