package uuid

import (
	"testing"
)

func TestNewV7(t *testing.T) {
	// Test default generation with fixed timestamp and random data
	idv7 := UUID{0x06, 0x0A, 0x11, 0xC2, 0xCE, 0xB7, 0x79, 0xA2, 0xB1, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD}
	testRand(0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF)
	CurrentTime = testTime()
	id, err := NewV7()
	if err != nil {
		t.Error("New failed to generate a UUIDv7")
	}
	if id != idv7 {
		t.Errorf("uuid.NewV7() = %v, want %v", id, idv7)
	}
	if id.Version() != 7 {
		t.Errorf("uuid.NewV7() generated UUID with wrong version %d", id.Version())
	}

	// Test generation with sequence counter
	UseSequenceCounter = true

	// First ID with sequence counter
	idv7 = UUID{0x06, 0x0A, 0x11, 0xC2, 0xCE, 0xB7, 0x79, 0xA2, 0xB1, 0x00, 0x00, 0x01, 0x23, 0x45, 0x67, 0x89}
	testRand(0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF)
	CurrentTime = testTime()
	id, err = NewV7()
	if err != nil {
		t.Error("New failed to generate a UUIDv7")
	}
	if id != idv7 {
		t.Errorf("uuid.NewV7() = %v, want %v", id, idv7)
	}
	if id.Version() != 7 {
		t.Errorf("uuid.NewV7() generated UUID with wrong version %d", id.Version())
	}

	// Second ID with sequence counter
	idv7 = UUID{0x06, 0x0A, 0x11, 0xC2, 0xCE, 0xB7, 0x79, 0xA2, 0xB1, 0x00, 0x01, 0x01, 0x23, 0x45, 0x67, 0x89}
	testRand(0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF)
	CurrentTime = testTime()
	id, err = NewV7()
	if err != nil {
		t.Error("New failed to generate a UUIDv7")
	}
	if id != idv7 {
		t.Errorf("uuid.NewV7() = %v, want %v", id, idv7)
	}
	if id.Version() != 7 {
		t.Errorf("uuid.NewV7() generated UUID with wrong version %d", id.Version())
	}

	// ID with sequence counter at different time
	idv7 = UUID{0x06, 0x0A, 0x11, 0xC2, 0xCE, 0xB7, 0x79, 0xA2, 0xB2, 0x00, 0x00, 0x01, 0x23, 0x45, 0x67, 0x89}
	testRand(0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF)
	CurrentTime = testTime(1)
	id, err = NewV7()
	if err != nil {
		t.Error("New failed to generate a UUIDv7")
	}
	if id != idv7 {
		t.Errorf("uuid.NewV7() = %v, want %v", id, idv7)
	}
	if id.Version() != 7 {
		t.Errorf("uuid.NewV7() generated UUID with wrong version %d", id.Version())
	}

	// Test default generation with random generator error
	UseSequenceCounter = false
	testRand()
	_, err = NewV7()
	if err == nil {
		t.Error("New did not fail when no random data was available")
	}

	// Test generation with sequence counter with random generator error
	UseSequenceCounter = true
	testRand()
	_, err = NewV7()
	if err == nil {
		t.Error("New did not fail when no random data was available")
	}
}
