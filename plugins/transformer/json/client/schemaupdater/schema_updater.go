package schemaupdater

import (
	"fmt"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/cloudquery/cloudquery/plugins/transformer/json/client/util"
)

// SchemaUpdater takes an `arrow.Schema` and knows how to make simple subsequent changes to it.
// It doesn't know which table it belongs to or if the changes make sense.
type SchemaUpdater struct {
	schema *arrow.Schema
}

func New(sc *arrow.Schema) *SchemaUpdater {
	return &SchemaUpdater{schema: sc}
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
				arrow.Field{Name: fmt.Sprintf("%s__%s", fieldName, subFieldName), Type: typ, Nullable: true},
			)
			if err != nil {
				return nil, err
			}
		}
	}

	return s.schema, nil
}

func typeFromString(t string) (arrow.DataType, error) {
	switch t {
	case "utf8":
		return arrow.BinaryTypes.String, nil
	case "int64":
		return arrow.PrimitiveTypes.Int64, nil
	case "timestamp":
		return arrow.FixedWidthTypes.Timestamp_us, nil
	case "bool":
		return arrow.FixedWidthTypes.Boolean, nil
	}
	return nil, fmt.Errorf("unknown type: %s", t)
}
