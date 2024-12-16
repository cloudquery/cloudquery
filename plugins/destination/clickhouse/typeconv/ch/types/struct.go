package types

import (
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
)

// https://clickhouse.com/docs/en/sql-reference/data-types/tuple
func structType(structType *arrow.StructType) (string, error) {
	defs, err := definitions(structType.Fields())
	if err != nil {
		return "", err
	}

	return "Tuple(" + strings.Join(defs, ", ") + ")", nil
}

func definitions(fields []arrow.Field) ([]string, error) {
	res := make([]string, len(fields))
	var err error

	for i, field := range fields {
		res[i], err = FieldDefinition(field)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}
