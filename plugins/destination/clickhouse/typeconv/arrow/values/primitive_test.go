package values

import (
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/float16"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"github.com/google/uuid"
)

func Test_primitive(t *testing.T) {
	for _, tc := range []testCase{
		{_type: new(arrow.BooleanType), value: true, expected: true},

		{_type: new(arrow.Uint8Type), value: uint8(123), expected: uint8(123)},
		{_type: new(arrow.Uint16Type), value: uint16(12345), expected: uint16(12345)},
		{_type: new(arrow.Uint32Type), value: uint32(1234567), expected: uint32(1234567)},
		{_type: new(arrow.Uint64Type), value: uint64(123456789), expected: uint64(123456789)},

		{_type: new(arrow.Int8Type), value: int8(-123), expected: int8(-123)},
		{_type: new(arrow.Int16Type), value: int16(-12345), expected: int16(-12345)},
		{_type: new(arrow.Int32Type), value: int32(-1234567), expected: int32(-1234567)},
		{_type: new(arrow.Int64Type), value: int64(-123456789), expected: int64(-123456789)},

		{_type: new(arrow.Float16Type), value: float32(-3.45), expected: float16.New(float32(-3.45))},
		{_type: new(arrow.Float32Type), value: float32(-3.4567), expected: float32(-3.4567)},
		{_type: new(arrow.Float64Type), value: float64(-3.456789), expected: float64(-3.456789)},

		{_type: new(arrow.StringType), value: "ABC", expected: "ABC"},
		{_type: types.NewUUIDType(), value: uuid.NameSpaceDNS, expected: uuid.NameSpaceDNS},
	} {
		ensureRecord(t, tc)
	}
}
