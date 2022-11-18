package recipes

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ecs/models"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ECSResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "clusters",
			Struct:      &types.Cluster{},
			Description: "https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Cluster.html",
			SkipFields:  []string{"ClusterArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ClusterArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"ClusterTasks()",
				"ClusterServices()",
				"ClusterContainerInstances()",
			},
		},
		{
			SubService:  "cluster_tasks",
			Struct:      &types.Task{},
			Description: "https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Task.html",
			SkipFields:  []string{"Tags", "TaskArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("TaskArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTags`,
					},
					{
						Name:     "task_protection",
						Type:     schema.TypeJSON,
						Resolver: `getEcsTaskProtection`,
					},
				}...),
			Relations: []string{},
		},
		{
			SubService:  "cluster_services",
			Struct:      &types.Service{},
			Description: "https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Service.html",
			SkipFields:  []string{"ServiceArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ServiceArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{},
		},
		{
			SubService:  "cluster_container_instances",
			Struct:      &types.ContainerInstance{},
			Description: "https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerInstance.html",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "cluster_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
			Relations: []string{},
		},
		{
			SubService:          "task_definitions",
			Struct:              &models.TaskDefinitionWrapper{},
			Description:         "https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskDefinition.html",
			SkipFields:          []string{"TaskDefinitionArn", "Tags"},
			PreResourceResolver: "getTaskDefinition",
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
						Resolver: `resolveEcsTaskDefinitionTags`,
					},
				}...),
			Relations: []string{},
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "ecs"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("ecs")`
		structName := reflect.ValueOf(r.Struct).Elem().Type().Name()
		if strings.Contains(structName, "Wrapper") {
			r.UnwrapEmbeddedStructs = true
		}
	}
	return resources
}
