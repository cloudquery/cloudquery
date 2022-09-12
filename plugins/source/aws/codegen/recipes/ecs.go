package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var ecsResources = []*Resource{
	{
		SubService:          "task_definitions",
		Struct:              &types.TaskDefinition{},
		SkipFields:          []string{"TaskDefinitionArn"},
		PreResourceResolver: "getEcsTaskDefinition",
		ExtraColumns: append(
			defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("TaskDefinitionArn")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "tags",
					Type:     schema.TypeJSON,
					Resolver: `schema.PathResolver("Tags")`,
				},
			}...),
	},
}

func EcsResources() []*Resource {
	for _, r := range ecsResources {
		r.Service = "ecs"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("ecs")`
	}
	return ecsResources
}
