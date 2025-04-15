package uuid

import (
	"bytes"
	"errors"
	"net"
	"reflect"
	"testing"
	"time"
)

const (
	testVecTimeCustom int64 = 1621171244987654321 // 2021-05-16T13:20:44.987654321Z
	testVecTimeRFC    int64 = 1645557742000000000 // 2022-02-22T07:22:22.00Z
)

func testPrepare(tm int64, rand []byte, randN uint32, macAddr net.HardwareAddr) {
	currentTime = func() time.Time { return time.Unix(0, tm) }
	randomSource = bytes.NewBuffer(rand).Read
	randomUint32 = func(uint32) uint32 { return randN }
	mac = macAddr
}

func TestSetMACAddress(t *testing.T) {
	tests := []struct {
		name    string
		mac     net.HardwareAddr
		wantErr bool
	}{
		{
			name:    "Valid MAC address",
			mac:     net.HardwareAddr{0x01, 0x02, 0x03, 0x04, 0x05, 0x06},
			wantErr: false,
		},
		{
			name:    "Invalid MAC address length (too short)",
			mac:     net.HardwareAddr{0x01, 0x02, 0x03, 0x04, 0x05},
			wantErr: true,
		},
		{
			name:    "Invalid MAC address length (too long)",
			mac:     net.HardwareAddr{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07},
			wantErr: true,
		},
		{
			name:    "Empty MAC address",
			mac:     net.HardwareAddr{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetMACAddress(tt.mac)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetMACAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If we didn't expect an error, verify the MAC address was set correctly
			if !tt.wantErr {
				if !reflect.DeepEqual(mac, tt.mac) {
					t.Errorf("SetMACAddress() did not set MAC correctly, got = %v, want %v", mac, tt.mac)
				}
			}
		})
	}
}

func TestUseHardwareMAC(t *testing.T) {
	tests := []struct {
		name       string
		interfaces []net.Interface
		interfErr  error
		wantErr    bool
		wantMAC    net.HardwareAddr
	}{
		{
			name: "Valid interface found",
			interfaces: []net.Interface{
				{
					HardwareAddr: net.HardwareAddr{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF},
				},
			},
			interfErr: nil,
			wantErr:   false,
			wantMAC:   net.HardwareAddr{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF},
		},
		{
			name: "Multiple interfaces, use first valid",
			interfaces: []net.Interface{
				{
					HardwareAddr: nil,
				},
				{
					HardwareAddr: net.HardwareAddr{0x01, 0x02},
				},
				{
					HardwareAddr: net.HardwareAddr{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF},
				},
			},
			interfErr: nil,
			wantErr:   false,
			wantMAC:   net.HardwareAddr{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF},
		},
		{
			name:       "No interfaces found",
			interfaces: []net.Interface{},
			interfErr:  nil,
			wantErr:    true,
			wantMAC:    nil,
		},
		{
			name:       "Error getting interfaces",
			interfaces: nil,
			interfErr:  errors.New("network error"),
			wantErr:    true,
			wantMAC:    nil,
		},
		{
			name: "No valid MAC addresses",
			interfaces: []net.Interface{
				{
					HardwareAddr: net.HardwareAddr{0x01, 0x02, 0x03},
				},
				{
					HardwareAddr: net.HardwareAddr{0x01, 0x02},
				},
			},
			interfErr: nil,
			wantErr:   true,
			wantMAC:   nil,
		},
	}

	// Save original function to restore later
	originalNetInterfaces := netInterfaces

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock netInterfaces function
			netInterfaces = func() ([]net.Interface, error) {
				return tt.interfaces, tt.interfErr
			}

			// Reset netInterfaces after test
			defer func() {
				netInterfaces = originalNetInterfaces
			}()

			err := UseHardwareMAC()
			if (err != nil) != tt.wantErr {
				t.Errorf("UseHardwareMAC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(mac, tt.wantMAC) {
				t.Errorf("UseHardwareMAC() did not set MAC correctly, got = %v, want %v", mac, tt.wantMAC)
			}
		})
	}
}

func TestUUID_IsNil(t *testing.T) {
	tests := []struct {
		name string
		UUID *UUID
		want bool
	}{
		{"NotNil", &UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x46, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, false},
		{"Nil", &UUID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.UUID.IsNil(); got != tt.want {
				t.Errorf("UUID.IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_IsMax(t *testing.T) {
	tests := []struct {
		name string
		UUID *UUID
		want bool
	}{
		{"NotMax", &UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x46, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, false},
		{"Max", &UUID{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.UUID.IsMax(); got != tt.want {
				t.Errorf("UUID.IsMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_Version(t *testing.T) {
	tests := []struct {
		name string
		UUID *UUID
		want int
	}{
		{"V0", &UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x06, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, 0},
		{"V1", &UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x16, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, 1},
		{"V2", &UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x26, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, 2},
		{"V3", &UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x36, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, 3},
		{"V4", &UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x46, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, 4},
		{"V5", &UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x56, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, 5},
		{"V6", &UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x66, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, 6},
		{"V7", &UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x76, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, 7},
		{"V8", &UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x86, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, 8},
		{"V9", &UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x96, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, 9},
		{"V10", &UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0xA6, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, 10},
		{"V11", &UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0xB6, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, 11},
		{"V12", &UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0xC6, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, 12},
		{"V13", &UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0xD6, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, 13},
		{"V14", &UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0xE6, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, 14},
		{"V15", &UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0xF6, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.UUID.Version(); got != tt.want {
				t.Errorf("UUID.Version() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_Timestamp(t *testing.T) {
	tests := []struct {
		name string
		uuid UUID
		want time.Time
	}{
		{"UUIDv7", UUID{0x01, 0x79, 0x75, 0x56, 0x0F, 0xBB, 0x70, 0x00, 0x81, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}, time.Unix(1621171244, 987*1000000)},
		{"UUIDv4", UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x46, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, time.Unix(0, 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.uuid.Timestamp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UUID.Timestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intervalsSinceEpoch(t *testing.T) {
	tests := []struct {
		name     string
		fakeTime int64
		want     int64
	}{
		{"TestTime", testVecTimeCustom, 138404640449876543},
		{"TestTime+1d", testVecTimeCustom + 86400000000000, 138404640449876543 + 864000000000},
		{"TestTime-3y", testVecTimeCustom - 94672800000000000, 138404640449876543 - 946728000000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testPrepare(tt.fakeTime, nil, 0, nil)
			diff1 := time.Date(1800, 1, 1, 0, 0, 0, 0, time.UTC).Sub(time.Date(1582, 10, 15, 0, 0, 0, 0, time.UTC)).Nanoseconds() / 100
			diff2 := currentTime().Sub(time.Date(1800, 1, 1, 0, 0, 0, 0, time.UTC)).Nanoseconds() / 100
			diff := diff1 + diff2
			if got := intervalsSinceEpoch(); got != tt.want {
				t.Errorf("intervalsSinceEpoch() = %v, want %v, calculated %v", got, tt.want, diff)
			}
		})
	}
}

func TestNamespaceDNS(t *testing.T) {
	expected := UUID{0x6B, 0xA7, 0xB8, 0x10, 0x9D, 0xAD, 0x11, 0xD1, 0x80, 0xB4, 0x00, 0xC0, 0x4F, 0xD4, 0x30, 0xC8}
	actual := NamespaceDNS()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("NamespaceDNS() = %v, want %v", actual, expected)
	}
}

func TestNamespaceURL(t *testing.T) {
	expected := UUID{0x6B, 0xA7, 0xB8, 0x11, 0x9D, 0xAD, 0x11, 0xD1, 0x80, 0xB4, 0x00, 0xC0, 0x4F, 0xD4, 0x30, 0xC8}
	actual := NamespaceURL()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("NamespaceURL() = %v, want %v", actual, expected)
	}
}

func TestNamespaceOID(t *testing.T) {
	expected := UUID{0x6B, 0xA7, 0xB8, 0x12, 0x9D, 0xAD, 0x11, 0xD1, 0x80, 0xB4, 0x00, 0xC0, 0x4F, 0xD4, 0x30, 0xC8}
	actual := NamespaceOID()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("NamespaceOID() = %v, want %v", actual, expected)
	}
}

func TestNamespaceX500(t *testing.T) {
	expected := UUID{0x6B, 0xA7, 0xB8, 0x14, 0x9D, 0xAD, 0x11, 0xD1, 0x80, 0xB4, 0x00, 0xC0, 0x4F, 0xD4, 0x30, 0xC8}
	actual := NamespaceX500()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("NamespaceX500() = %v, want %v", actual, expected)
	}
}
