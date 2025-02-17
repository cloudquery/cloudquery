package client

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
)

type CustomUnit64 uint64

// See https://stackoverflow.com/questions/60520865/will-mongodb-go-driver-automatically-convert-uint64-to-bson-type-overflows-er
// Please note we don't need the decoder part from the above link, as we already cast back to uint64 when we pass the MongoDB saved value to arrow Uint64Builder
func getRegistry() *bsoncodec.Registry {
	reg := bson.NewRegistry()

	customUnit64 := reflect.TypeOf(CustomUnit64(0))
	reg.RegisterTypeEncoder(
		customUnit64,
		bsoncodec.ValueEncoderFunc(func(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
			if !val.IsValid() || val.Type() != customUnit64 {
				return bsoncodec.ValueEncoderError{
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
