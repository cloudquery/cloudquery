package client

import (
	"encoding/hex"

	"github.com/jackc/pgtype"
)

type UUID struct {
	pgtype.UUID
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
