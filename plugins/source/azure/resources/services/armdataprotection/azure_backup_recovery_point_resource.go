// Code generated by codegen; DO NOT EDIT.

package armdataprotection

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AzureBackupRecoveryPointResource() *schema.Table {
	return &schema.Table{
		Name:      "azure_armdataprotection_azure_backup_recovery_point_resource",
		Resolver:  fetchAzureBackupRecoveryPointResource,
		Multiplex: client.SubscriptionResourceGroupMultiplex,
		Columns: []schema.Column{
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

func fetchAzureBackupRecoveryPointResource(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armdataprotection.NewRecoveryPointsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
