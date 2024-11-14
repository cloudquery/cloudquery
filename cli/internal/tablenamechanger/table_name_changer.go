package tablenamechanger

import (
	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

// TableNameChanger manages table name changes due to transformers, on a per-destination basis.
// It can learn about table name changes and apply those changes on further steps.
type TableNameChanger struct {
	tableNameChanges map[string]map[string]string
}

func New(destinationSpecs []specs.Destination) *TableNameChanger {
	tnc := &TableNameChanger{
		tableNameChanges: make(map[string]map[string]string),
	}
	for _, destinationSpec := range destinationSpecs {
		tnc.tableNameChanges[destinationSpec.Name] = make(map[string]string)
	}
	return tnc
}

func (c TableNameChanger) UpdateTableNames(destinationName string, tables map[string]bool) map[string]bool {
	newTables := make(map[string]bool, len(tables))
	for oldTableName := range tables {
		if newTableName, ok := c.tableNameChanges[destinationName][oldTableName]; ok {
			delete(tables, oldTableName)
			newTables[newTableName] = true
		} else {
			newTables[oldTableName] = true
		}
	}
	return newTables
}

func (c TableNameChanger) UpdateTableNamesSlice(destinationName string, tables []string) []string {
	newTables := make([]string, len(tables))
	for i, oldTableName := range tables {
		newTables[i] = c.UpdateTableName(destinationName, oldTableName)
	}
	return newTables
}

func (c TableNameChanger) UpdateTableName(destinationName string, oldTableName string) string {
	if newTableName, ok := c.tableNameChanges[destinationName][oldTableName]; ok {
		return newTableName
	}
	return oldTableName
}

func (c *TableNameChanger) LearnTableNameChange(destinationName, oldTableName string, schemaBytes []byte) error {
	// Implicit assumption that table name changes are deterministic.
	// Therefore, once we learned about a table name change, we don't need to learn it again.
	if c.tableNameChanges[destinationName][oldTableName] != "" {
		return nil
	}

	sc, err := plugin.NewSchemaFromBytes(schemaBytes)
	if err != nil {
		return err
	}
	newTableName, ok := sc.Metadata().GetValue(schema.MetadataTableName)
	if !ok {
		return nil // this would be an error, but let it fail at a more relevant step
	}
	c.tableNameChanges[destinationName][oldTableName] = newTableName
	return nil
}
