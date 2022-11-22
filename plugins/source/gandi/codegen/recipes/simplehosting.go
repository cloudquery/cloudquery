package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/go-gandi/go-gandi/simplehosting"
)

func SimpleHostingResources() []Resource {
	return []Resource{
		{
			DataStruct:   &simplehosting.Instance{},
			PKColumns:    []string{"id"},
			ExtraColumns: []codegen.ColumnDefinition{SharingIDColumn},
			Relations:    []string{"InstanceVhosts()"},
			TableName:    "gandi_simplehosting_instances",

			Template:         "resource_manual",
			Package:          "simplehosting",
			TableFuncName:    "Instances",
			Filename:         "instances.go",
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
			TableName: "gandi_simplehosting_instance_vhosts",

			Template:         "resource_manual",
			Package:          "simplehosting",
			TableFuncName:    "InstanceVhosts",
			Filename:         "instance_vhosts.go",
			ResolverFuncName: "fetchInstanceVhosts",
		},
	}
}
