package drift

import (
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

type traversedTable struct {
	*schema.Table
	Parent *traversedTable
}

func (t *traversedTable) Resolvers(name string, builtin bool) []string {
	colNames := make([]string, 0, len(t.Table.Columns))
	for _, c := range t.Table.Columns {
		m := c.Meta()
		if m == nil || m.Resolver == nil || m.Resolver.Name != name || m.Resolver.Builtin != builtin {
			continue
		}
		colNames = append(colNames, c.Name)
	}
	return colNames
}

func (t *traversedTable) AccountIDColumn() string {
	cols := t.Resolvers("github.com/cloudquery/cq-provider-aws/client.ResolveAWSAccount", false)
	if len(cols) == 0 {
		return ""
	}
	return cols[0]
}

func (t *traversedTable) ParentIDColumn() string {
	cols := t.Resolvers("schema.ParentIdResolver", true)
	if len(cols) == 0 {
		return ""
	}
	return cols[0]
}

// AutoIgnoreColumns only returns columns that are cq-specific and should be ignored when drifting, such as parent resolvers, regions or account IDs
func (t *traversedTable) AutoIgnoreColumns() []string {
	var colNames []string
	for _, c := range t.Table.Columns {
		m := c.Meta()
		if m == nil || m.Resolver == nil {
			continue
		}
		switch m.Resolver.Name {
		case "schema.ParentIdResolver",
			"github.com/cloudquery/cq-provider-aws/client.ResolveAWSAccount",
			"github.com/cloudquery/cq-provider-aws/client.ResolveAWSRegion":
			colNames = append(colNames, c.Name)
		}
	}
	return colNames
}

// NonCQColumns returns a list of columns with autoignore applied
func (t *traversedTable) NonCQColumns() []string {
	ig := t.AutoIgnoreColumns()
	igm := make(map[string]struct{}, len(ig))
	for i := range ig {
		igm[ig[i]] = struct{}{}
	}

	cols := make([]string, 0, len(t.Table.Columns))
	for _, c := range t.Table.Columns {
		if _, ok := igm[c.Name]; ok {
			continue
		}
		cols = append(cols, c.Name)
	}

	return cols
}

// NonCQPrimaryKeys returns primary keys with autoignore applied
func (t *traversedTable) NonCQPrimaryKeys() []string {
	ig := t.AutoIgnoreColumns()
	igm := make(map[string]struct{}, len(ig))
	for i := range ig {
		igm[ig[i]] = struct{}{}
	}

	pks := make([]string, 0, len(t.Table.Options.PrimaryKeys))
	for _, pk := range t.Table.Options.PrimaryKeys {
		if _, ok := igm[pk]; ok {
			continue
		}
		pks = append(pks, pk)
	}

	return pks
}

// traverseResourceTable iterates each resource and sets up parent relationships, returning a traversedTable map with parents set.
// On the topmost level resources are accessible with both their resource ID ("ec2.instances") and their table name ("aws_ec2_instances")
// Since child resources don't have resource IDs, they are only accessed by table name
func traverseResourceTable(resources map[string]*schema.Table) map[string]*traversedTable {
	tableMap := make(map[string]*traversedTable)

	var setTableMap func(res *schema.Table, parent *traversedTable)
	setTableMap = func(res *schema.Table, parent *traversedTable) {
		tableMap[res.Name] = &traversedTable{
			Table:  res,
			Parent: parent,
		}
		for _, rel := range res.Relations {
			setTableMap(rel, tableMap[res.Name])
		}
	}

	for resId, res := range resources {
		tableMap[resId] = &traversedTable{Table: res}
		setTableMap(res, nil)
	}

	return tableMap
}
