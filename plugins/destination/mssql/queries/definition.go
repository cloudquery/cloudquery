package queries

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
)

type Definition struct {
	Name    string
	typ     string
	notNull bool
	unique  bool
}

// sanitized returns definition copy with name sanitized
func (d *Definition) sanitized() *Definition {
	r := *d
	r.Name = sanitizeID(r.Name)
	return &r
}

func (d *Definition) Type() string {
	return d.typ
}

func (d *Definition) Constraint() string {
	var res []string

	if d.unique {
		res = append(res, "UNIQUE")
	}

	if d.notNull {
		res = append(res, "NOT NULL")
	}

	return strings.Join(res, " ")
}

// Nullable returns definition copy that will allow nullable values
func (d *Definition) Nullable() *Definition {
	return &Definition{
		Name: d.Name,
		typ:  d.typ,
	}
}

func NewDefinition(name, typ string, nullable bool) *Definition {
	d := &Definition{
		Name:    strings.ToLower(name),
		typ:     strings.ToLower(typ),
		notNull: !nullable,
	}
	// add unique for _cq_id
	d.unique = d.Name == schema.CqIDColumn.Name
	return d
}

func GetDefinition(column *schema.Column, pkEnabled bool) *Definition {
	def := &Definition{
		Name: column.Name,
		typ:  SQLType(column.Type),
	}

	switch {
	case column.Name == schema.CqIDColumn.Name:
		// _cq_id column should always have a "UNIQUE NOT NULL" constraint
		def.unique = true
		def.notNull = true
	case pkEnabled && column.CreationOptions.PrimaryKey:
		def.notNull = true
	}

	return def
}

type Definitions []*Definition

func (defs Definitions) Get(name string) *Definition {
	for _, d := range defs {
		if d.Name == name {
			return d
		}
	}
	return nil
}

// GetDefinitions returns sanitized Definitions
func GetDefinitions(columns schema.ColumnList, pkEnabled bool) Definitions {
	definitions := make(Definitions, len(columns))

	for i, col := range columns {
		definitions[i] = GetDefinition(&col, pkEnabled).sanitized()
	}

	return definitions
}
