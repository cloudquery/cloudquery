// Code generated by codegen; DO NOT EDIT.

package armmariadb

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func PerformanceTierProperties() *schema.Table {
	return &schema.Table{
		Name:      "azure_armmariadb_performance_tier_properties",
		Resolver:  fetchPerformanceTierProperties,
		Multiplex: client.SubscriptionResourceGroupMultiplex,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "max_backup_retention_days",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxBackupRetentionDays"),
			},
			{
				Name:     "max_large_storage_mb",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxLargeStorageMB"),
			},
			{
				Name:     "max_storage_mb",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxStorageMB"),
			},
			{
				Name:     "min_backup_retention_days",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinBackupRetentionDays"),
			},
			{
				Name:     "min_large_storage_mb",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinLargeStorageMB"),
			},
			{
				Name:     "min_storage_mb",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinStorageMB"),
			},
			{
				Name:     "service_level_objectives",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ServiceLevelObjectives"),
			},
		},
	}
}

func fetchPerformanceTierProperties(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armmariadb.NewServerBasedPerformanceTierClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
