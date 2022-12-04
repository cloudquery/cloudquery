// Code generated by codegen; DO NOT EDIT.

package armrecoveryservices

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ReplicationUsage() *schema.Table {
	return &schema.Table{
		Name:      "azure_armrecoveryservices_replication_usage",
		Resolver:  fetchReplicationUsage,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "jobs_summary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("JobsSummary"),
			},
			{
				Name:     "monitoring_summary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MonitoringSummary"),
			},
			{
				Name:     "protected_item_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ProtectedItemCount"),
			},
			{
				Name:     "recovery_plan_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("RecoveryPlanCount"),
			},
			{
				Name:     "recovery_services_provider_auth_type",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("RecoveryServicesProviderAuthType"),
			},
			{
				Name:     "registered_servers_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("RegisteredServersCount"),
			},
		},
	}
}

func fetchReplicationUsage(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armrecoveryservices.NewReplicationUsagesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
