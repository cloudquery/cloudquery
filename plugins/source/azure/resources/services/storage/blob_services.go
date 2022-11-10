// Auto generated code - DO NOT EDIT.

package storage

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
)

func blobServices() *schema.Table {
	return &schema.Table{
		Name:        "azure_storage_blob_services",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage#BlobServiceProperties`,
		Resolver:    fetchStorageBlobServices,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "storage_account_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "cors",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Cors"),
			},
			{
				Name:     "default_service_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultServiceVersion"),
			},
			{
				Name:     "delete_retention_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DeleteRetentionPolicy"),
			},
			{
				Name:     "is_versioning_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsVersioningEnabled"),
			},
			{
				Name:     "automatic_snapshot_policy_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AutomaticSnapshotPolicyEnabled"),
			},
			{
				Name:     "change_feed",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ChangeFeed"),
			},
			{
				Name:     "restore_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RestorePolicy"),
			},
			{
				Name:     "container_delete_retention_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ContainerDeleteRetentionPolicy"),
			},
			{
				Name:     "last_access_time_tracking_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LastAccessTimeTrackingPolicy"),
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
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

func fetchStorageBlobServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Storage.BlobServices

	account := parent.Item.(storage.Account)
	if !isBlobSupported(&account) {
		return nil
	}

	resource, err := client.ParseResourceID(*account.ID)
	if err != nil {
		return err
	}
	response, err := svc.List(ctx, resource.ResourceGroup, *account.Name)
	if err != nil {
		return err
	}
	if response.Value == nil {
		return nil
	}
	res <- *response.Value

	return nil
}
