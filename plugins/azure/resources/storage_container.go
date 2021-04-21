package resources

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func storageContainers() *schema.Table {
	return &schema.Table{
		Name:         "azure_storage_containers",
		Resolver:     fetchStorageContainers,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "account_id",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
			},
			{
				Name:     "version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContainerProperties.Version"),
			},
			{
				Name:     "deleted",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ContainerProperties.Deleted"),
			},
			{
				Name:     "deleted_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ContainerProperties.DeletedTime.Time"),
			},
			{
				Name:     "remaining_retention_days",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ContainerProperties.RemainingRetentionDays"),
			},
			{
				Name:     "default_encryption_scope",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContainerProperties.DefaultEncryptionScope"),
			},
			{
				Name:     "deny_encryption_scope_override",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ContainerProperties.DenyEncryptionScopeOverride"),
			},
			{
				Name:     "public_access",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContainerProperties.PublicAccess"),
			},
			{
				Name:     "last_modified_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ContainerProperties.LastModifiedTime.Time"),
			},
			{
				Name:     "lease_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContainerProperties.LeaseStatus"),
			},
			{
				Name:     "lease_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContainerProperties.LeaseState"),
			},
			{
				Name:     "lease_duration",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContainerProperties.LeaseDuration"),
			},
			{
				Name:     "metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ContainerProperties.Metadata"),
			},
			{
				Name:     "immutability_policy",
				Type:     schema.TypeJSON,
				Resolver: resolveImmutabilityPolicy,
			},
			{
				Name:     "legal_hold",
				Type:     schema.TypeJSON,
				Resolver: resolveLegalHold,
			},
			{
				Name:     "has_legal_hold",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ContainerProperties.HasLegalHold"),
			},
			{
				Name:     "has_immutability_policy",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ContainerProperties.HasImmutabilityPolicy"),
			},
			{
				Name: "etag",
				Type: schema.TypeString,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchStorageContainers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Storage.Containers
	acc := parent.Item.(storage.Account)
	resourceDetails, err := client.ParseResourceID(*acc.ID)
	if err != nil {
		return err
	}
	response, err := svc.List(ctx, resourceDetails.ResourceGroup, *acc.Name, "", "", "")
	if err != nil {
		return err
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil

}
func resolveImmutabilityPolicy(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	container := resource.Item.(storage.ListContainerItem)
	if container.ImmutabilityPolicy == nil || container.ImmutabilityPolicy.UpdateHistory == nil {
		return nil
	}
	data, err := container.ImmutabilityPolicy.MarshalJSON()
	if err != nil {
		return err
	}
	resource.Set("immutability_policy", data)
	return nil
}

func resolveLegalHold(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	container := resource.Item.(storage.ListContainerItem)

	data, err := container.LegalHold.MarshalJSON()
	if err != nil {
		return err
	}
	resource.Set("legal_hold", data)
	return nil
}
