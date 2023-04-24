package values

import (
	"math/rand"
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/decimal128"
	"github.com/apache/arrow/go/v12/arrow/decimal256"
	"github.com/apache/arrow/go/v12/arrow/memory"
	"github.com/stretchr/testify/require"
)

func Test_decimal128Value(t *testing.T) {
	const (
		N     = 100
		valid = 30
	)
	values := make([]float64, N)
	for i := range values {
		values[i] = rand.Float64()*1000 + rand.Float64()
	}

	for _, _type := range []*arrow.Decimal128Type{
		{Precision: 10, Scale: 0},
		{Precision: 10, Scale: 3},
		{Precision: 10, Scale: 5},
		{Precision: 15, Scale: 0},
		{Precision: 15, Scale: 3},
		{Precision: 15, Scale: 5},
		{Precision: 15, Scale: 10},
		{Precision: 23, Scale: 15},
		{Precision: 32, Scale: 2},
		{Precision: 32, Scale: 15},
		{Precision: 38, Scale: 32},
	} {
		t.Run(_type.String(), func(t *testing.T) {
			in := make([]*decimal128.Num, len(values))
			bld := array.NewDecimal128Builder(memory.DefaultAllocator, _type)
			for i, val := range values {
				num, err := decimal128.FromFloat64(val, _type.Precision, _type.Scale)
				require.NoError(t, err)

				trimmed := num.ToFloat64(_type.Scale)
				num, err = decimal128.FromFloat64(trimmed, _type.Precision, _type.Scale)
				require.NoError(t, err)
				bld.Append(num)
				in[i] = &num
			}

			arr := bld.NewArray()
			out := decimalValue[decimal128.Num](arr.(*array.Decimal128))

			require.Equal(t, N, len(out))
			for i, out := range out {
				require.NotNil(t, out)

				// trim to valid symbols
				in := in[i]
				inS, outS := []rune(in.ToString(_type.Scale)), []rune(out.StringFixed(_type.Scale))
				require.Exactly(t, len(inS), len(outS))
				if l := len(inS); l > valid {
					inS, outS = inS[:valid], outS[:valid]
				}

				require.Exactly(t, string(inS), string(outS))
			}
		})
	}
}

func Test_decimal256Value(t *testing.T) {
	const (
		N     = 100
		valid = 30
	)

	values := make([]float64, N)
	for i := range values {
		values[i] = rand.Float64()*1000 + rand.Float64()
	}

	for _, _type := range []*arrow.Decimal256Type{
		{Precision: 10, Scale: 0},
		{Precision: 10, Scale: 3},
		{Precision: 10, Scale: 5},
		{Precision: 15, Scale: 0},
		{Precision: 15, Scale: 3},
		{Precision: 15, Scale: 5},
		{Precision: 15, Scale: 10},
		{Precision: 23, Scale: 15},
		{Precision: 32, Scale: 2},
		{Precision: 32, Scale: 15},
		{Precision: 38, Scale: 32},
	} {
		t.Run(_type.String(), func(t *testing.T) {
			in := make([]*decimal256.Num, len(values))
			bld := array.NewDecimal256Builder(memory.DefaultAllocator, _type)
			for i, val := range values {
				num, err := decimal256.FromFloat64(val, _type.Precision, _type.Scale)
				require.NoError(t, err)

				trimmed := num.ToFloat64(_type.Scale)
				num, err = decimal256.FromFloat64(trimmed, _type.Precision, _type.Scale)
				require.NoError(t, err)
				bld.Append(num)
				in[i] = &num
			}

			arr := bld.NewArray()
			out := decimalValue[decimal256.Num](arr.(*array.Decimal256))

			require.Equal(t, N, len(out))
			for i, out := range out {
				require.NotNil(t, out)

				// trim to valid symbols
				in := in[i]
				require.NotNil(t, in)
				inS, outS := []rune(in.ToString(_type.Scale)), []rune(out.StringFixed(_type.Scale))
				require.Exactly(t, len(inS), len(outS))
				if l := len(inS); l > valid {
					inS, outS = inS[:valid], outS[:valid]
				}

				require.Exactly(t, string(inS), string(outS))
			}
		})
	}
}
