package client

import (
	"reflect"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/oracle/oci-go-sdk/v65/common"
)

func OracleTypeTransformer(field reflect.StructField) (arrow.DataType, error) {
	fieldType := field.Type

	if fieldType.Kind() == reflect.Ptr {
		fieldType = fieldType.Elem()
	}

	if fieldType.Kind() == reflect.Struct && fieldType == reflect.TypeOf(common.SDKTime{}) {
		return arrow.FixedWidthTypes.Timestamp_us, nil
	}

	return transformers.DefaultTypeTransformer(field)
}
