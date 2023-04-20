package typeconv

import (
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func CanonizedSchemas(scs schema.Schemas) (schema.Schemas, error) {
	schemas := make(schema.Schemas, len(scs))
	var err error
	for i, sc := range scs {
		schemas[i], err = CanonizedSchema(sc)
		if err != nil {
			return nil, err
		}
	}
	return schemas, nil
}

func CanonizedSchema(sc *arrow.Schema) (*arrow.Schema, error) {
	fields := make([]arrow.Field, len(sc.Fields()))
	for i, fld := range sc.Fields() {
		canonized, err := CanonizedField(fld)
		if err != nil {
			return nil, err
		}
		fields[i] = *canonized
	}

	metadata := sc.Metadata()
	return arrow.NewSchema(fields, &metadata), nil
}

// CanonizedField allows to know what type we associate the Apache Arrow type with.
// Several different Apache Arrow types will produce the same canonical type & that'll be the type we'll use in the database.
func CanonizedField(field arrow.Field) (*arrow.Field, error) {
	// 1 - convert to the ClickHouse
	chType, err := ClickHouseDefinitions(field)
	if err != nil {
		return nil, err
	}

	// 2 - convert back to Apache Arrow
	return ArrowField(field.Name, chType[0])
}
