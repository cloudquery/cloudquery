package values

import (
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/decimal128"
	"github.com/apache/arrow/go/v16/arrow/decimal256"
	"github.com/shopspring/decimal"
)

func buildDecimal128(builder *array.Decimal128Builder, value any) {
	v, ok := unwrap[decimal.Decimal](value)
	if !ok {
		builder.AppendNull()
		return
	}
	builder.Append(decimal128.FromBigInt(v.Coefficient()))
}

func buildDecimal256(builder *array.Decimal256Builder, value any) {
	v, ok := unwrap[decimal.Decimal](value)
	if !ok {
		builder.AppendNull()
		return
	}
	builder.Append(decimal256.FromBigInt(v.Coefficient()))
}
