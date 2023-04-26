package queries

import (
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
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

func GetDefinition(field arrow.Field, pkEnabled bool) *Definition {
	def := &Definition{
		Name:    field.Name,
		typ:     SQLType(field.Type),
		notNull: !field.Nullable,
		unique:  schema.IsUnique(field),
	}

	if pkEnabled && schema.IsPk(field) {
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
func GetDefinitions(sc *arrow.Schema, pkEnabled bool) Definitions {
	definitions := make(Definitions, len(sc.Fields()))

	for i, field := range sc.Fields() {
		definitions[i] = GetDefinition(field, pkEnabled).sanitized()
	}

	return definitions
}
