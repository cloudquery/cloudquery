package client

import (
	"reflect"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CustomUnit64 uint64

// See https://stackoverflow.com/questions/60520865/will-mongodb-go-driver-automatically-convert-uint64-to-bson-type-overflows-er
// Please note we don't need the decoder part from the above link, as we already cast back to uint64 when we pass the MongoDB saved value to arrow Uint64Builder
func getRegistry() *bson.Registry {
	reg := bson.NewRegistry()

	customUnit64 := reflect.TypeOf(CustomUnit64(0))
	reg.RegisterTypeEncoder(
		customUnit64,
		bson.ValueEncoderFunc(func(_ bson.EncodeContext, vw bson.ValueWriter, val reflect.Value) error {
			if !val.IsValid() || val.Type() != customUnit64 {
				return bson.ValueEncoderError{
					Name:     "CustomUnit64EncodedValue",
					Types:    []reflect.Type{customUnit64},
					Received: val,
				}
			}
			return vw.WriteInt64(int64(val.Uint()))
		}),
	)

	return reg
}
