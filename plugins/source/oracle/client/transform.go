package client

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/oracle/oci-go-sdk/v65/common"
)

func OracleTypeTransformer(field reflect.StructField) (schema.ValueType, error) {
	fieldType := field.Type

	if fieldType.Kind() == reflect.Ptr {
		fieldType = fieldType.Elem()
	}

	if fieldType.Kind() == reflect.Struct && fieldType == reflect.TypeOf(common.SDKTime{}) {
		return schema.TypeTimestamp, nil
	}

	return transformers.DefaultTypeTransformer(field)
}
