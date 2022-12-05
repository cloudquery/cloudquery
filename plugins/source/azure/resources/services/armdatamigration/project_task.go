// Code generated by codegen; DO NOT EDIT.

package armdatamigration

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ProjectTask() *schema.Table {
	return &schema.Table{
		Name:      "azure_armdatamigration_project_task",
		Resolver:  fetchProjectTask,
		Multiplex: client.SubscriptionResourceGroupMultiplex,
		Columns: []schema.Column{
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "system_data",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SystemData"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}

func fetchProjectTask(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armdatamigration.NewTasksClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(cl.ResourceGroup, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
