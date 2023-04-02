package uuid

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name     string
		args     args
		wantUUID UUID
		wantErr  bool
	}{
		{"Normal", args{"686e7778-f9f0-4622-a13e-c2441ce4ae41"}, UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x46, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, false},
		{"UpperCase", args{"686E7778-F9F0-4622-A13E-C2441CE4AE41"}, UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x46, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, false},
		{"MixedCase", args{"686E7778-F9f0-4622-A13e-C2441cE4aE41"}, UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x46, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, false},
		{"WrongLength", args{"01234567-89ab-cdef"}, UUID{}, true},
		{"WrongDashPlacement", args{"012345678-9ab-cdef-012345-6789abcdef"}, UUID{}, true},
		{"NonHexCharactersSection1", args{"ghijklmn-abcd-abcd-abcd-0123456789ab"}, UUID{}, true},
		{"NonHexCharactersSection2", args{"abcdef01-opqr-abcd-abcd-0123456789ab"}, UUID{}, true},
		{"NonHexCharactersSection3", args{"abcdef01-abcd-stuv-abcd-0123456789ab"}, UUID{}, true},
		{"NonHexCharactersSection4", args{"abcdef01-abcd-abcd-wxyz-0123456789ab"}, UUID{}, true},
		{"NonHexCharactersSection5", args{"abcdef01-abcd-abcd-abcd-ghijklmnopqr"}, UUID{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUUID, err := Parse(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUUID, tt.wantUUID) {
				t.Errorf("Parse() = %v, want %v", gotUUID, tt.wantUUID)
			}
		})
	}
}

func TestParseBytes(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name     string
		args     args
		wantUUID UUID
		wantErr  bool
	}{
		{"Normal", args{[]byte{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x46, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}}, UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x46, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, false},
		{"WrongLength", args{[]byte{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x46, 0x22, 0xa1, 0x3e}}, UUID{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUUID, err := ParseBytes(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUUID, tt.wantUUID) {
				t.Errorf("ParseBytes() = %v, want %v", gotUUID, tt.wantUUID)
			}
		})
	}
}

func TestUUID_String(t *testing.T) {
	tests := []struct {
		name    string
		UUID    *UUID
		wantOut string
	}{
		{"ConvertToString", &UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x46, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, "686e7778-f9f0-4622-a13e-c2441ce4ae41"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := tt.UUID.String(); gotOut != tt.wantOut {
				t.Errorf("UUID.String() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestUUID_MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		uuid    UUID
		want    []byte
		wantErr bool
	}{
		{"Normal", UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x46, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, []byte("686e7778-f9f0-4622-a13e-c2441ce4ae41"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uuid.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("UUID.MarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UUID.MarshalText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_UnmarshalText(t *testing.T) {
	type args struct {
		in []byte
	}
	tests := []struct {
		name    string
		uuid    *UUID
		args    args
		wantErr bool
	}{
		{"Normal", &UUID{}, args{[]byte{0x36, 0x38, 0x36, 0x65, 0x37, 0x37, 0x37, 0x38, 0x2d, 0x66, 0x39, 0x66, 0x30, 0x2d, 0x34, 0x36, 0x32, 0x32, 0x2d, 0x61, 0x31, 0x33, 0x65, 0x2d, 0x63, 0x32, 0x34, 0x34, 0x31, 0x63, 0x65, 0x34, 0x61, 0x65, 0x34, 0x31}}, false},
		{"InvalidID", &UUID{}, args{[]byte{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.uuid.UnmarshalText(tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("UUID.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUUID_MarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		uuid    UUID
		want    []byte
		wantErr bool
	}{
		{"Normal", UUID{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x46, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, []byte{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x46, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uuid.MarshalBinary()
			if (err != nil) != tt.wantErr {
				t.Errorf("UUID.MarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UUID.MarshalBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_UnmarshalBinary(t *testing.T) {
	type args struct {
		in []byte
	}
	tests := []struct {
		name    string
		uuid    *UUID
		args    args
		wantErr bool
	}{
		{"Normal", &UUID{}, args{[]byte{0x68, 0x6e, 0x77, 0x78, 0xf9, 0xf0, 0x46, 0x22, 0xa1, 0x3e, 0xc2, 0x44, 0x1c, 0xe4, 0xae, 0x41}}, false},
		{"InvalidID", &UUID{}, args{[]byte{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.uuid.UnmarshalBinary(tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("UUID.UnmarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
