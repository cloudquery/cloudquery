package client

import (
	"bytes"
	"database/sql/driver"
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/jackc/pgtype"
)

// Code came from: https://github.com/jackc/pgtype/blob/75446032b914bb0be5e07da29c976034c0a666cf/uuid.go#L3-L230

type UUID struct {
	Bytes  [16]byte
	Status pgtype.Status
}

func (dst *UUID) Set(src interface{}) error {
	if src == nil {
		*dst = UUID{Status: pgtype.Null}
		return nil
	}

	if value, ok := src.(interface{ Get() interface{} }); ok {
		value2 := value.Get()
		if value2 != value {
			return dst.Set(value2)
		}
	}

	switch value := src.(type) {
	case [16]byte:
		*dst = UUID{Bytes: value, Status: pgtype.Present}
	case []byte:
		if value != nil {
			if len(value) != 16 {
				return fmt.Errorf("[]byte must be 16 bytes to convert to UUID: %d", len(value))
			}
			*dst = UUID{Status: pgtype.Present}
			copy(dst.Bytes[:], value)
		} else {
			*dst = UUID{Status: pgtype.Null}
		}
	case string:
		uuid, err := parseUUID(value)
		if err != nil {
			return err
		}
		*dst = UUID{Bytes: uuid, Status: pgtype.Present}
	case *string:
		if value == nil {
			*dst = UUID{Status: pgtype.Null}
		} else {
			return dst.Set(*value)
		}
	default:
		if originalSrc, ok := underlyingUUIDType(src); ok {
			return dst.Set(originalSrc)
		}
		return fmt.Errorf("cannot convert %v to UUID", value)
	}

	return nil
}

func (dst UUID) Get() interface{} {
	switch dst.Status {
	case pgtype.Present:
		// CQ-Change: Return entire object, not just Bytes
		return dst
	case pgtype.Null:
		return nil
	default:
		return dst.Status
	}
}

func (src *UUID) AssignTo(dst interface{}) error {
	switch src.Status {
	case pgtype.Present:
		switch v := dst.(type) {
		case *[16]byte:
			*v = src.Bytes
			return nil
		case *[]byte:
			*v = make([]byte, 16)
			copy(*v, src.Bytes[:])
			return nil
		case *string:
			*v = encodeUUID(src.Bytes)
			return nil
		default:
			if nextDst, retry := pgtype.GetAssignToDstType(v); retry {
				return src.AssignTo(nextDst)
			}
		}
	case pgtype.Null:
		return pgtype.NullAssignTo(dst)
	}

	return fmt.Errorf("cannot assign %v into %T", src, dst)
}

// parseUUID converts a string UUID in standard form to a byte array.
func parseUUID(src string) (dst [16]byte, err error) {
	switch len(src) {
	case 36:
		src = src[0:8] + src[9:13] + src[14:18] + src[19:23] + src[24:]
	case 32:
		// dashes already stripped, assume valid
	default:
		// assume invalid.
		return dst, fmt.Errorf("cannot parse UUID %v", src)
	}

	buf, err := hex.DecodeString(src)
	if err != nil {
		return dst, err
	}

	copy(dst[:], buf)
	return dst, err
}

// encodeUUID converts a uuid byte array to UUID standard string form.
func encodeUUID(src [16]byte) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", src[0:4], src[4:6], src[6:8], src[8:10], src[10:16])
}

func (dst *UUID) DecodeText(ci *pgtype.ConnInfo, src []byte) error {
	if src == nil {
		*dst = UUID{Status: pgtype.Null}
		return nil
	}

	if len(src) != 36 {
		return fmt.Errorf("invalid length for UUID: %v", len(src))
	}

	buf, err := parseUUID(string(src))
	if err != nil {
		return err
	}

	*dst = UUID{Bytes: buf, Status: pgtype.Present}
	return nil
}

func (dst *UUID) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {
	if src == nil {
		*dst = UUID{Status: pgtype.Null}
		return nil
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for UUID: %v", len(src))
	}

	*dst = UUID{Status: pgtype.Present}
	copy(dst.Bytes[:], src)
	return nil
}

func (src UUID) EncodeText(ci *pgtype.ConnInfo, buf []byte) ([]byte, error) {
	switch src.Status {
	case pgtype.Null:
		return nil, nil
	case pgtype.Undefined:
		return nil, errUndefined
	}

	return append(buf, encodeUUID(src.Bytes)...), nil
}

func (src UUID) EncodeBinary(ci *pgtype.ConnInfo, buf []byte) ([]byte, error) {
	switch src.Status {
	case pgtype.Null:
		return nil, nil
	case pgtype.Undefined:
		return nil, errUndefined
	}

	return append(buf, src.Bytes[:]...), nil
}

// Scan implements the database/sql Scanner interface.
func (dst *UUID) Scan(src interface{}) error {
	if src == nil {
		*dst = UUID{Status: pgtype.Null}
		return nil
	}

	switch src := src.(type) {
	case string:
		return dst.DecodeText(nil, []byte(src))
	case []byte:
		srcCopy := make([]byte, len(src))
		copy(srcCopy, src)
		return dst.DecodeText(nil, srcCopy)
	}

	return fmt.Errorf("cannot scan %T", src)
}

// Value implements the database/sql/driver Valuer interface.
func (src UUID) Value() (driver.Value, error) {
	return pgtype.EncodeValueText(src)
}

func (src UUID) MarshalJSON() ([]byte, error) {
	switch src.Status {
	case pgtype.Present:
		var buff bytes.Buffer
		buff.WriteByte('"')
		buff.WriteString(encodeUUID(src.Bytes))
		buff.WriteByte('"')
		return buff.Bytes(), nil
	case pgtype.Null:
		return []byte("null"), nil
	case pgtype.Undefined:
		return nil, errUndefined
	}
	return nil, errBadStatus
}

func (dst *UUID) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		return dst.Set(nil)
	}
	if len(src) != 38 {
		return fmt.Errorf("invalid length for UUID: %v", len(src))
	}
	return dst.Set(string(src[1 : len(src)-1]))
}

func underlyingUUIDType(val interface{}) (interface{}, bool) {
	refVal := reflect.ValueOf(val)

	if refVal.Kind() == reflect.Ptr {
		if refVal.IsNil() {
			return time.Time{}, false
		}
		convVal := refVal.Elem().Interface()
		return convVal, true
	}

	uuidType := reflect.TypeOf([16]byte{})
	if refVal.Type().ConvertibleTo(uuidType) {
		return refVal.Convert(uuidType).Interface(), true
	}

	return nil, false
}

var errUndefined = errors.New("cannot encode status undefined")
var errBadStatus = errors.New("invalid status")

// CQ-Change: Overload string to enable printing of struct
func (u UUID) String() string {
	buf := make([]byte, 36)

	hex.Encode(buf[0:8], u.Bytes[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], u.Bytes[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], u.Bytes[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], u.Bytes[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], u.Bytes[10:])
	return string(buf)
}
