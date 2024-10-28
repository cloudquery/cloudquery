package schemaupdater

import (
	"fmt"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/cloudquery/cloudquery/plugins/transformer/jsonflattener/client/util"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

const (
	BoolType      = "bool"
	UTF8Type      = "utf8"
	Int64Type     = "int64"
	Float64Type   = "float64"
	JSONType      = "json"
	TimestampType = "timestamp[us, tz=UTC]"
)

// SchemaUpdater takes an `arrow.Schema` and knows how to make simple subsequent changes to it.
// It doesn't know which table it belongs to or if the changes make sense.
type SchemaUpdater struct {
	schema      *arrow.Schema
	uniqueNames map[string]bool
}

func New(sc *arrow.Schema) *SchemaUpdater {
	return &SchemaUpdater{schema: sc, uniqueNames: make(map[string]bool)}
}

func (s *SchemaUpdater) AddJSONFlattenedFields(fieldTypeSchemas map[string]map[string]string) (*arrow.Schema, error) {
	fieldNames := util.SortedKeys(fieldTypeSchemas)

	for _, fieldName := range fieldNames {
		typeSchema := fieldTypeSchemas[fieldName]
		subFieldNames := util.SortedKeys(typeSchema)

		for _, subFieldName := range subFieldNames {
			subFieldType := typeSchema[subFieldName]
			typ, err := typeFromString(subFieldType)
			if err != nil {
				return nil, err
			}
			s.schema, err = s.schema.AddField(
				s.schema.NumFields(),
				arrow.Field{Name: s.uniqueName(fieldName, subFieldName, ""), Type: typ, Nullable: true},
			)
			if err != nil {
				return nil, err
			}
		}
	}

	return s.schema, nil
}

func (s *SchemaUpdater) uniqueName(fieldName, subFieldName, suffix string) string {
	candidateName := fmt.Sprintf("%s__%s%s", fieldName, subFieldName, suffix)
	if s.uniqueNames[candidateName] {
		return s.uniqueName(fieldName, subFieldName, fmt.Sprintf("%s_", suffix))
	}
	s.uniqueNames[candidateName] = true
	return candidateName
}

func typeFromString(t string) (arrow.DataType, error) {
	switch t {
	case UTF8Type:
		return arrow.BinaryTypes.String, nil
	case Int64Type:
		return arrow.PrimitiveTypes.Int64, nil
	case TimestampType:
		return arrow.FixedWidthTypes.Timestamp_us, nil
	case BoolType:
		return arrow.FixedWidthTypes.Boolean, nil
	case Float64Type:
		return arrow.PrimitiveTypes.Float64, nil
	case JSONType:
		return types.NewJSONType(), nil
	}
	return nil, fmt.Errorf("unknown type: %s", t)
}
