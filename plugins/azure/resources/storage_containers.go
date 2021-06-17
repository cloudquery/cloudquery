package resources

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func StorageContainers() *schema.Table {
	return &schema.Table{
		Name:         "azure_storage_containers",
		Description:  "Azure storage container",
		Resolver:     fetchStorageContainers,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "account_id",
				Description: "Azure storage account id",
				Type:        schema.TypeUUID,
				Resolver:    schema.ParentIdResolver,
			},
			{
				Name:        "version",
				Description: "The version of the deleted blob container",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContainerProperties.Version"),
			},
			{
				Name:        "deleted",
				Description: "Indicates whether the blob container was deleted",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ContainerProperties.Deleted"),
			},
			{
				Name:        "deleted_time",
				Description: "Blob container deletion time",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("ContainerProperties.DeletedTime.Time"),
			},
			{
				Name:        "remaining_retention_days",
				Description: "Remaining retention days for soft deleted blob container",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ContainerProperties.RemainingRetentionDays"),
			},
			{
				Name:        "default_encryption_scope",
				Description: "Default the container to use specified encryption scope for all writes",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContainerProperties.DefaultEncryptionScope"),
			},
			{
				Name:        "deny_encryption_scope_override",
				Description: "Block override of encryption scope from the container default",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ContainerProperties.DenyEncryptionScopeOverride"),
			},
			{
				Name:        "public_access",
				Description: "Specifies whether data in the container may be accessed publicly and the level of access Possible values include: 'PublicAccessContainer', 'PublicAccessBlob', 'PublicAccessNone'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContainerProperties.PublicAccess"),
			},
			{
				Name:        "last_modified_time",
				Description: "Returns the date and time the container was last modified",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("ContainerProperties.LastModifiedTime.Time"),
			},
			{
				Name:        "lease_status",
				Description: "The lease status of the container Possible values include: 'LeaseStatusLocked', 'LeaseStatusUnlocked'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContainerProperties.LeaseStatus"),
			},
			{
				Name:        "lease_state",
				Description: "Lease state of the container Possible values include: 'LeaseStateAvailable', 'LeaseStateLeased', 'LeaseStateExpired', 'LeaseStateBreaking', 'LeaseStateBroken'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContainerProperties.LeaseState"),
			},
			{
				Name:        "lease_duration",
				Description: "Specifies whether the lease on a container is of infinite or fixed duration, only when the container is leased Possible values include: 'Infinite', 'Fixed'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContainerProperties.LeaseDuration"),
			},
			{
				Name:        "metadata",
				Description: "A name-value pair to associate with the container as metadata",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ContainerProperties.Metadata"),
			},
			{
				Name:        "immutability_policy",
				Description: "The ImmutabilityPolicy property of the container",
				Type:        schema.TypeJSON,
				Resolver:    resolveStorageContainerImmutabilityPolicy,
			},
			{
				Name:        "legal_hold",
				Description: "The LegalHold property of the container",
				Type:        schema.TypeJSON,
				Resolver:    resolveStorageContainerLegalHold,
			},
			{
				Name:        "has_legal_hold",
				Description: "The hasLegalHold public property is set to true by SRP if there are at least one existing tag The hasLegalHold public property is set to false by SRP if all existing legal hold tags are cleared out There can be a maximum of 1000 blob containers with hasLegalHold=true for a given account",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ContainerProperties.HasLegalHold"),
			},
			{
				Name:        "has_immutability_policy",
				Description: "The hasImmutabilityPolicy public property is set to true by SRP if ImmutabilityPolicy has been created for this container The hasImmutabilityPolicy public property is set to false by SRP if ImmutabilityPolicy has not been created for this container",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ContainerProperties.HasImmutabilityPolicy"),
			},
			{
				Name:        "etag",
				Description: "Resource Etag",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_id",
				Description: "Fully qualified resource ID for the resource Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of the resource Eg \"MicrosoftCompute/virtualMachines\" or \"MicrosoftStorage/storageAccounts\"",
				Type:        schema.TypeString,
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

func resolveStorageContainerImmutabilityPolicy(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	container := resource.Item.(storage.ListContainerItem)
	if container.ImmutabilityPolicy == nil || container.ImmutabilityPolicy.UpdateHistory == nil {
		return nil
	}
	data, err := container.ImmutabilityPolicy.MarshalJSON()
	if err != nil {
		return err
	}
	return resource.Set("immutability_policy", data)
}

func resolveStorageContainerLegalHold(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	container := resource.Item.(storage.ListContainerItem)

	data, err := container.LegalHold.MarshalJSON()
	if err != nil {
		return err
	}
	return resource.Set("legal_hold", data)
}
