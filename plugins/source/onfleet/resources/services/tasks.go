package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/onfleet/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/keplr-team/go-onfleet/onfleet"
)

func Tasks() *schema.Table {
	return &schema.Table{
		Name:        "onfleet_tasks",
		Resolver:    fetchTasks,
		Description: "https://docs.onfleet.com/reference/list-tasks",
		Transform: transformers.TransformWithStruct(&onfleet.Task{},
			client.OnfleetSharedTransformers()...),
		Columns: []schema.Column{
			{
				Name:     "organization_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveOrganizationId,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
