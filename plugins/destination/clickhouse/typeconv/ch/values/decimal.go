package values

import (
	"math/big"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/shopspring/decimal"
)

type decimalToBigInt interface {
	BigInt() *big.Int
}

func decimalValue[A decimalToBigInt, ARR primitive[A]](arr ARR) []*decimal.Decimal {
	// decimal.Decimal = big.Int * (10 ^ exp)
	scale := arr.DataType().(arrow.DecimalType).GetScale()

	res := make([]*decimal.Decimal, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) {
			val := decimal.NewFromBigInt(arr.Value(i).BigInt(), -scale)
			res[i] = &val
		}
	}
	return res
}
