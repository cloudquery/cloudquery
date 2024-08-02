package schemaupdater

import (
	"bytes"
	"errors"
	"text/template"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

// SchemaUpdater takes an `arrow.Schema` and knows how to make simple subsequent changes to it.
// It doesn't know which table it belongs to or if the changes make sense.
type SchemaUpdater struct {
	schema *arrow.Schema
}

func New(sc *arrow.Schema) *SchemaUpdater {
	return &SchemaUpdater{schema: sc}
}

func (s *SchemaUpdater) RemoveColumnIndices(colIndices map[int]struct{}) *arrow.Schema {
	oldMetadata := s.schema.Metadata()
	oldFields := s.schema.Fields()
	newFields := make([]arrow.Field, 0, len(oldFields)-len(colIndices))
	for i := range oldFields {
		if _, ok := colIndices[i]; ok {
			continue
		}
		newFields = append(newFields, oldFields[i])
	}
	s.schema = arrow.NewSchema(newFields, &oldMetadata)
	return s.schema
}

func (s *SchemaUpdater) AddStringColumnAtPos(columnName string, zeroIndexedPosition int, isNullable bool) (*arrow.Schema, error) {
	if zeroIndexedPosition == -1 {
		zeroIndexedPosition = s.schema.NumFields()
	}
	return s.schema.AddField(
		zeroIndexedPosition,
		arrow.Field{Name: columnName, Type: arrow.BinaryTypes.String, Nullable: isNullable},
	)
}

func (s *SchemaUpdater) ChangeTableName(newTableNamePattern string) (*arrow.Schema, error) {
	existingMetadata := s.schema.Metadata()
	tableName, ok := existingMetadata.GetValue(schema.MetadataTableName)
	if !ok {
		return nil, errors.New("table name not found in record's metadata")
	}

	type tpl struct {
		OldName string
	}

	t, err := template.New("table_name").Parse(newTableNamePattern)
	if err != nil {
		return nil, err
	}

	var tplBuf bytes.Buffer
	if err := t.Execute(&tplBuf, tpl{OldName: tableName}); err != nil {
		return nil, err
	}

	newName := tplBuf.String()

	m := existingMetadata.ToMap()
	m[schema.MetadataTableName] = newName
	newMetadata := arrow.MetadataFrom(m)
	return arrow.NewSchema(s.schema.Fields(), &newMetadata), nil
}
