package resource

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/gertd/go-pluralize"
)

var (
	SubscriptionIDCol = codegen.ColumnDefinition{
		Name:        "subscription_id",
		Type:        schema.TypeString,
		Resolver:    "client.SubscriptionIDResolver",
		Description: "Azure subscription ID",
	}
)

func (r *Resource) propagateColumns() {
	namePresent := make(map[string]bool)
	for _, col := range r.Table.Columns {
		namePresent[col.Name] = true
	}

	prefixes := []string{"properties_"}
	columns := make(codegen.ColumnDefinitions, 0, len(r.Table.Columns))
	for _, col := range r.Table.Columns {
		for _, pfx := range prefixes {
			if strings.HasPrefix(col.Name, pfx) {
				newName := strings.TrimPrefix(col.Name, pfx)
				if !namePresent[newName] {
					col.Name = newName
				}
				break
			}
		}
		namePresent[col.Name] = true
		columns = append(columns, col)
	}

	if r.parent != nil && r.parent.hasField("ID") {
		parentIDCol := r.parentIDCol()
		if namePresent[parentIDCol.Name] {
			parentIDCol.Name = "parent_" + parentIDCol.Name
		}
		columns = append(columns, parentIDCol)
	}

	r.Table.Columns = columns
}

func (r *Resource) parentIDCol() codegen.ColumnDefinition {
	return codegen.ColumnDefinition{
		Name:          pluralize.NewClient().Singular(r.parent.SubService) + "_id",
		Type:          schema.TypeString,
		Resolver:      `schema.ParentColumnResolver("id")`,
		IgnoreInTests: false,
		Options:       schema.ColumnCreationOptions{},
	}
}
