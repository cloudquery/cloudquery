package schemaupdater

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
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

func (s *SchemaUpdater) AddTimestampColumnAtPos(columnName string, zeroIndexedPosition int, isNullable bool) (*arrow.Schema, error) {
	if zeroIndexedPosition == -1 {
		zeroIndexedPosition = s.schema.NumFields()
	}
	return s.schema.AddField(
		zeroIndexedPosition,
		arrow.Field{Name: columnName, Type: arrow.FixedWidthTypes.Timestamp_us, Nullable: isNullable},
	)
}

func (s *SchemaUpdater) RenameColumn(oldName, newName string) (*arrow.Schema, error) {
	oldFields := s.schema.Fields()

	newFields := make([]arrow.Field, len(oldFields))

	for i, f := range oldFields {
		if f.Name == oldName {
			f.Name = newName
		}
		newFields[i] = f
	}
	metadata := s.schema.Metadata()
	return arrow.NewSchema(newFields, &metadata), nil
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

func (s *SchemaUpdater) AddPrimaryKeys(newPks []string) (*arrow.Schema, error) {
	table, err := schema.NewTableFromArrowSchema(s.schema)
	if err != nil {
		return nil, err
	}
	for _, newPk := range newPks {
		newCol := table.Columns.Get(newPk)
		if newCol == nil {
			return nil, fmt.Errorf("new primary key column: %s not found in: %s", newCol.Name, table.Name)
		}
		newCol.PrimaryKey = true
	}
	return table.ToArrowSchema(), nil
}
