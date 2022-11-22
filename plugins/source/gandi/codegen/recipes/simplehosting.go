package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/go-gandi/go-gandi/simplehosting"
)

func SimpleHostingResources() []Resource {
	return []Resource{
		{
			DataStruct:       &simplehosting.Instance{},
			PKColumns:        []string{"id"},
			ExtraColumns:     []codegen.ColumnDefinition{SharingIDColumn},
			Relations:        []string{"InstanceVhosts()"},
			Template:         "resource_manual",
			TableName:        "gandi_simplehosting_instances",
			TableFuncName:    "Instances",
			Filename:         "instances.go",
			Package:          "simplehosting",
			ResolverFuncName: "fetchInstances",
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
			Template:         "resource_manual",
			TableName:        "gandi_simplehosting_instance_vhosts",
			TableFuncName:    "InstanceVhosts",
			Filename:         "instance_vhosts.go",
			Package:          "simplehosting",
			ResolverFuncName: "fetchInstanceVhosts",
		},
	}
}
