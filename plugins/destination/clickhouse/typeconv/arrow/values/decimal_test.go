package values

import (
	"math/rand"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/decimal128"
	"github.com/apache/arrow-go/v18/arrow/decimal256"
	"github.com/shopspring/decimal"
)

func Test_decimal128(t *testing.T) {
	value := decimal.NewFromFloat((rand.Float64()-rand.Float64())*1000 + rand.Float64())
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
		ensureRecord(t, testCase{
			dataType: decimalType,
			value:    &value,
			expected: decimal128.FromBigInt(value.Coefficient()),
		})
	}
}

func Test_decimal256(t *testing.T) {
	value := decimal.NewFromFloat((rand.Float64()-rand.Float64())*1000 + rand.Float64())
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
		ensureRecord(t, testCase{
			dataType: decimalType,
			value:    &value,
			expected: decimal256.FromBigInt(value.Coefficient()),
		})
	}
}
