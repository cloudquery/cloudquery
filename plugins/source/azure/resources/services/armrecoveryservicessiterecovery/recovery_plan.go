// Code generated by codegen; DO NOT EDIT.

package armrecoveryservicessiterecovery

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RecoveryPlan() *schema.Table {
	return &schema.Table{
		Name:      "azure_armrecoveryservicessiterecovery_recovery_plan",
		Resolver:  fetchRecoveryPlan,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties"),
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
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}

func fetchRecoveryPlan(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armrecoveryservicessiterecovery.NewReplicationRecoveryPlansClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
