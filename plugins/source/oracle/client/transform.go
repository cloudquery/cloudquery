package client

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
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

	return codegen.DefaultTypeTransformer(field)
}
