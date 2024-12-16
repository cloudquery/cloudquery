package values

import (
	"math/rand"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/decimal128"
	"github.com/apache/arrow-go/v18/arrow/decimal256"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/require"
)

func Test_decimal128Value(t *testing.T) {
	const amount = 100
	values := make([]float64, amount)
	for i := range values {
		values[i] = (rand.Float64()-rand.Float64())*1000 + rand.Float64()
	}

	for _, decimalType := range []*arrow.Decimal128Type{
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
		t.Run(decimalType.String(), func(t *testing.T) {
			in := make([]*decimal128.Num, len(values))
			builder := array.NewDecimal128Builder(memory.DefaultAllocator, decimalType)
			for i, val := range values {
				num, err := decimal128.FromFloat64(val, decimalType.Precision, decimalType.Scale)
				require.NoError(t, err)

				trimmed := num.ToFloat64(decimalType.Scale)
				num, err = decimal128.FromFloat64(trimmed, decimalType.Precision, decimalType.Scale)
				require.NoError(t, err)
				builder.Append(num)
				in[i] = &num
			}

			out := decimalValue(builder.NewDecimal128Array())

			require.Equal(t, amount, len(out))
			for i, out := range out {
				require.NotNil(t, out)

				in := in[i]
				out := decimal128.FromBigInt(out.Coefficient())
				require.Exactly(t, in.ToString(decimalType.Scale), out.ToString(decimalType.Scale))
			}
		})
	}
}

func Test_decimal256Value(t *testing.T) {
	const amount = 100

	values := make([]float64, amount)
	for i := range values {
		values[i] = (rand.Float64()-rand.Float64())*1000 + rand.Float64()
	}

	for _, decimalType := range []*arrow.Decimal256Type{
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
		t.Run(decimalType.String(), func(t *testing.T) {
			in := make([]*decimal256.Num, len(values))
			builder := array.NewDecimal256Builder(memory.DefaultAllocator, decimalType)
			for i, val := range values {
				num, err := decimal256.FromFloat64(val, decimalType.Precision, decimalType.Scale)
				require.NoError(t, err)

				trimmed := num.ToFloat64(decimalType.Scale)
				num, err = decimal256.FromFloat64(trimmed, decimalType.Precision, decimalType.Scale)
				require.NoError(t, err)
				builder.Append(num)
				in[i] = &num
			}

			out := decimalValue(builder.NewDecimal256Array())

			require.Equal(t, amount, len(out))
			for i, out := range out {
				require.NotNil(t, out)

				in := in[i]
				out := decimal256.FromBigInt(out.Coefficient())
				require.Exactly(t, in.ToString(decimalType.Scale), out.ToString(decimalType.Scale))
			}
		})
	}
}
