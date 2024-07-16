package transformations

import (
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v16/arrow"
)

// Transformation represents a transformation that can be applied to a recordBatch.
//
// - A transformation can transform both schema and data.
// - To transform the data, supply an SQLFn that returns the SQL query to apply to the recordBatch.
// - To transform the schema, supply a SchemaChangeFn that returns the new schema after applying the transformation.
// - If you change the schema in SQL but don't report it in the schema change function (or do it wrong), the sync will fail at runtime.
type Transformation struct {
	SQLFn          func(old *arrow.Schema) (string, error)
	SchemaChangeFn func(old *arrow.Schema) (*arrow.Schema, error)
}

func NewCustomDataOnlyTransformation(sql string) (Transformation, error) {
	return Transformation{
		SQLFn:          func(*arrow.Schema) (string, error) { return sql, nil },
		SchemaChangeFn: func(old *arrow.Schema) (*arrow.Schema, error) { return old, nil },
	}, nil
}

func NewCustomTransformation(sql string, schema *arrow.Schema) (Transformation, error) {
	return Transformation{
		SQLFn: func(*arrow.Schema) (string, error) { return sql, nil },
		SchemaChangeFn: func(old *arrow.Schema) (*arrow.Schema, error) {
			return schemaWithMergedMetadataFromOldSchema(old, schema), nil
		},
	}, nil
}

func NewAddFieldTransformation(sqlSelectFieldName, sqlSelectFieldSQL string, field arrow.Field) (Transformation, error) {
	return Transformation{
		SQLFn: func(*arrow.Schema) (string, error) {
			return fmt.Sprintf("SELECT *, (%v) AS %v FROM source_table", sqlSelectFieldSQL, sqlSelectFieldName), nil
		},
		SchemaChangeFn: func(old *arrow.Schema) (*arrow.Schema, error) {
			if old == nil {
				return nil, fmt.Errorf("old schema is nil")
			}
			newSchema, err := old.AddField(old.NumFields(), field)
			return newSchema, err
		},
	}, nil
}

func NewRemoveFieldTransformation(removeFieldName string) (Transformation, error) {
	sqlFn := func(schema *arrow.Schema) (string, error) {
		fieldNames := []string{}
		for _, field := range schema.Fields() {
			if field.Name == removeFieldName {
				continue
			}
			fieldNames = append(fieldNames, field.Name)
		}
		return fmt.Sprintf("SELECT %v FROM source_table", strings.Join(fieldNames, ", ")), nil
	}

	schemaChangeFn := func(old *arrow.Schema) (*arrow.Schema, error) {
		return cloneSchemaWithoutFields(old, removeFieldName), nil
	}

	return Transformation{
		SQLFn:          sqlFn,
		SchemaChangeFn: schemaChangeFn,
	}, nil
}

func NewUpdateFieldTransformation(updateFieldName string, sql string) (Transformation, error) {
	sqlFn := func(schema *arrow.Schema) (string, error) {
		fieldNames := []string{}
		for _, field := range schema.Fields() {
			fieldNames = append(fieldNames, field.Name)
		}
		sqlFields := []string{}
		for _, fieldName := range fieldNames {
			if fieldName == updateFieldName {
				sqlFields = append(sqlFields, fmt.Sprintf("(%v) AS %v", sql, updateFieldName))
			} else {
				sqlFields = append(sqlFields, fieldName)
			}
		}
		return fmt.Sprintf("SELECT %v FROM source_table", strings.Join(sqlFields, ", ")), nil
	}

	return Transformation{
		SQLFn:          sqlFn,
		SchemaChangeFn: func(old *arrow.Schema) (*arrow.Schema, error) { return old, nil },
	}, nil
}

func cloneSchemaWithoutFields(old *arrow.Schema, removeFieldNames ...string) *arrow.Schema {
	removeFieldNamesMap := make(map[string]struct{}, len(removeFieldNames))
	for _, removeFieldName := range removeFieldNames {
		removeFieldNamesMap[removeFieldName] = struct{}{}
	}

	newSchemaFields := make([]arrow.Field, 0, len(old.Fields())-1)
	for _, field := range old.Fields() {
		if _, ok := removeFieldNamesMap[field.Name]; ok {
			continue
		}
		newSchemaFields = append(newSchemaFields, field)
	}
	clonedMetadata := arrow.NewMetadata(
		append([]string{}, old.Metadata().Keys()...),
		append([]string{}, old.Metadata().Values()...),
	)
	return arrow.NewSchema(newSchemaFields, &clonedMetadata)
}

func schemaWithMergedMetadataFromOldSchema(old *arrow.Schema, new *arrow.Schema) *arrow.Schema {
	var (
		oldKeys    = old.Metadata().Keys()
		oldValues  = old.Metadata().Values()
		newKeys    = new.Metadata().Keys()
		newValues  = new.Metadata().Values()
		newKeysMap = make(map[string]struct{}, len(newKeys))
	)
	for i := range newKeys {
		newKeysMap[newKeys[i]] = struct{}{}
	}

	for i := range oldKeys {
		if _, ok := newKeysMap[oldKeys[i]]; !ok {
			newKeys = append(newKeys, oldKeys[i])
			newValues = append(newValues, oldValues[i])
		}
	}
	clonedMetadata := arrow.NewMetadata(newKeys, newValues)
	return arrow.NewSchema(new.Fields(), &clonedMetadata)
}
