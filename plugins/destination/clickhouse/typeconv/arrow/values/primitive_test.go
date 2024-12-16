package values

import (
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/float16"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/google/uuid"
)

func Test_primitive(t *testing.T) {
	for _, tc := range []testCase{
		{dataType: new(arrow.BooleanType), value: true, expected: true},

		{dataType: new(arrow.Uint8Type), value: uint8(123), expected: uint8(123)},
		{dataType: new(arrow.Uint16Type), value: uint16(12345), expected: uint16(12345)},
		{dataType: new(arrow.Uint32Type), value: uint32(1234567), expected: uint32(1234567)},
		{dataType: new(arrow.Uint64Type), value: uint64(123456789), expected: uint64(123456789)},

		{dataType: &arrow.Time32Type{Unit: arrow.Second}, value: time.Unix(123, 0).UTC(), expected: arrow.Time32(123)},
		{dataType: &arrow.Time32Type{Unit: arrow.Millisecond}, value: time.UnixMilli(123).UTC(), expected: arrow.Time32(123)},
		{dataType: &arrow.Time64Type{Unit: arrow.Microsecond}, value: time.UnixMicro(123).UTC(), expected: arrow.Time64(123)},
		{dataType: &arrow.Time64Type{Unit: arrow.Nanosecond}, value: time.Unix(0, 123).UTC(), expected: arrow.Time64(123)},

		{dataType: new(arrow.Int8Type), value: int8(-123), expected: int8(-123)},
		{dataType: new(arrow.Int16Type), value: int16(-12345), expected: int16(-12345)},
		{dataType: new(arrow.Int32Type), value: int32(-1234567), expected: int32(-1234567)},
		{dataType: new(arrow.Int64Type), value: int64(-123456789), expected: int64(-123456789)},

		{dataType: new(arrow.Float16Type), value: float32(-3.45), expected: float16.New(float32(-3.45))},
		{dataType: new(arrow.Float32Type), value: float32(-3.4567), expected: float32(-3.4567)},
		{dataType: new(arrow.Float64Type), value: float64(-3.456789), expected: float64(-3.456789)},

		{dataType: new(arrow.StringType), value: "ABC", expected: "ABC"},
		{dataType: types.NewUUIDType(), value: uuid.NameSpaceDNS, expected: uuid.NameSpaceDNS},
	} {
		ensureRecord(t, tc)
	}
}
