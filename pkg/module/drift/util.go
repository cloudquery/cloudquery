package drift

import (
	"fmt"

	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func getIACProvider(provs []*cqproto.GetProviderSchemaResponse) (*cqproto.GetProviderSchemaResponse, error) {
	var iacProv *cqproto.GetProviderSchemaResponse
	for _, p := range provs {
		if p.Name == string(iacTerraform) {
			if iacProv != nil {
				return nil, fmt.Errorf("only single IAC provider is supported at a time")
			}
			iacProv = p
		}
	}
	if iacProv == nil {
		return nil, fmt.Errorf("no IAC provider detected, can't continue")
	}

	return iacProv, nil
}

type traversedTable struct {
	*schema.Table
	Parent *traversedTable
}

func (t *traversedTable) Resolvers(name string, builtin bool) []string {
	var colNames []string
	for _, c := range t.Table.Columns {
		m := c.Meta()
		//if m != nil && m.Resolver != nil {
		//	fmt.Println("===", m.Resolver.Name)
		//}
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
