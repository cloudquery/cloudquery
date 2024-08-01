package schemaupdater

import "github.com/apache/arrow/go/v17/arrow"

// SchemaUpdater takes an `arrow.Schema` and knows how to make simple subsequent changes to it.
// It doesn't know which table it belongs to or if the changes make sense.
type SchemaUpdater struct {
	schema *arrow.Schema
}

func New(schema *arrow.Schema) *SchemaUpdater {
	return &SchemaUpdater{schema: schema}
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
