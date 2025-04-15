package uuid

import "testing"

func BenchmarkV1(b *testing.B) {
	for b.Loop() {
		NewV1()
	}
}

func BenchmarkV3(b *testing.B) {
	for b.Loop() {
		NewV3(NamespaceDNS(), "example.com")
	}
}

func BenchmarkV4(b *testing.B) {
	for b.Loop() {
		NewV4()
	}
}

func BenchmarkV5(b *testing.B) {
	for b.Loop() {
		NewV5(NamespaceDNS(), "example.com")
	}
}

func BenchmarkV6(b *testing.B) {
	for b.Loop() {
		NewV6()
	}
}

func BenchmarkV7(b *testing.B) {
	for b.Loop() {
		NewV7()
	}
}

func BenchmarkV8(b *testing.B) {
	for b.Loop() {
		NewV8([]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF})
	}
}
