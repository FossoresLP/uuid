package uuid

import (
	"database/sql/driver"
	"encoding/hex"
	"fmt"
)

// Parse parses a string as a UUID returning either the resulting UUID or an error
func Parse(str string) (UUID, error) {
	if len(str) != 36 {
		return UUID{}, fmt.Errorf("invalid length for UUID: %d", len(str))
	}
	if str[8] != '-' || str[13] != '-' || str[18] != '-' || str[23] != '-' {
		return UUID{}, fmt.Errorf("UUID format invalid")
	}
	uuid := UUID{}
	in := []byte(str)
	_, err := hex.Decode(uuid[:4], in[:8])
	if err != nil {
		return UUID{}, fmt.Errorf("UUID did contain unexpected character in segment %d", 1)
	}
	_, err = hex.Decode(uuid[4:6], in[9:13])
	if err != nil {
		return UUID{}, fmt.Errorf("UUID did contain unexpected character in segment %d", 2)
	}
	_, err = hex.Decode(uuid[6:8], in[14:18])
	if err != nil {
		return UUID{}, fmt.Errorf("UUID did contain unexpected character in segment %d", 3)
	}
	_, err = hex.Decode(uuid[8:10], in[19:23])
	if err != nil {
		return UUID{}, fmt.Errorf("UUID did contain unexpected character in segment %d", 4)
	}
	_, err = hex.Decode(uuid[10:16], in[24:36])
	if err != nil {
		return UUID{}, fmt.Errorf("UUID did contain unexpected character in segment %d", 5)
	}
	return uuid, nil
}

// ParseBytes parses a byte slice and returns the contained BINARY UUID or an error
func ParseBytes(bytes []byte) (uuid UUID, err error) {
	if len(bytes) != 16 {
		return uuid, fmt.Errorf("invalid length for binary UUID: %d", len(bytes))
	}
	copy(uuid[:], bytes)
	return
}

// ToString returns the string representation of a UUID
func (uuid UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

// MarshalText provides encoding.TextMarshaler
func (uuid UUID) MarshalText() ([]byte, error) {
	return []byte(uuid.String()), nil
}

// UnmarshalText provides encoding.TextUnmarshaler
func (uuid *UUID) UnmarshalText(in []byte) error {
	id, err := Parse(string(in))
	if err != nil {
		return err
	}
	*uuid = id
	return nil
}

// MarshalBinary provides encoding.BinaryMarshaler
func (uuid UUID) MarshalBinary() ([]byte, error) {
	return uuid[:], nil
}

// UnmarshalBinary provides encoding.BinaryUnmarshaler
func (uuid *UUID) UnmarshalBinary(in []byte) error {
	if len(in) != 16 {
		return fmt.Errorf("invalid length for binary UUID: %d", len(in))
	}
	copy(uuid[:], in)
	return nil
}

// Scan provides database/sql.Scanner
func (uuid *UUID) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		if len(v) != 16 {
			return fmt.Errorf("invalid length for binary UUID: %d", len(v))
		}
		copy(uuid[:], v)
	case string:
		id, err := Parse(v)
		if err != nil {
			return err
		}
		*uuid = id
	default:
		return fmt.Errorf("unknown type %t", v)
	}
	return nil
}

// Value provides database/sql/driver.Valuer
func (uuid UUID) Value() (driver.Value, error) {
	return uuid[:], nil
}
