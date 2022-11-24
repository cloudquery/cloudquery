package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/go-gandi/go-gandi/simplehosting"
)

func SimpleHostingResources() []*Resource {
	return []*Resource{
		{
			DataStruct:   &simplehosting.Instance{},
			PKColumns:    []string{"id"},
			ExtraColumns: []codegen.ColumnDefinition{SharingIDColumn},
			Relations:    []string{"SimplehostingInstanceVhosts()"},
			TableName:    "simplehosting_instances",
		},
		{
			DataStruct: &simplehosting.Vhost{},
			PKColumns:  []string{"instance_id", "fqdn"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "instance_id",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("id")`,
				},
			},
			TableName: "simplehosting_instance_vhosts",
		},
	}
}
