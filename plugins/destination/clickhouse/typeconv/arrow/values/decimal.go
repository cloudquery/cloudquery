package values

import (
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/decimal128"
	"github.com/apache/arrow/go/v12/arrow/decimal256"
	"github.com/shopspring/decimal"
)

func buildDecimal128(builder *array.Decimal128Builder, value *decimal.Decimal) {
	if value == (*decimal.Decimal)(nil) {
		builder.AppendNull()
		return
	}
	builder.Append(decimal128.FromBigInt((*value).Coefficient()))
}

func buildDecimal256(builder *array.Decimal256Builder, value *decimal.Decimal) {
	if value == (*decimal.Decimal)(nil) {
		builder.AppendNull()
		return
	}
	builder.Append(decimal256.FromBigInt((*value).Coefficient()))
}
