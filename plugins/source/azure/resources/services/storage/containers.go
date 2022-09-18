// Auto generated code - DO NOT EDIT.

package storage

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
)

func containers() *schema.Table {
	return &schema.Table{
		Name:     "azure_storage_containers",
		Resolver: fetchStorageContainers,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "cq_id_parent",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
			},
			{
				Name:     "version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Version"),
			},
			{
				Name:     "deleted",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Deleted"),
			},
			{
				Name:     "deleted_time",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DeletedTime"),
			},
			{
				Name:     "remaining_retention_days",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("RemainingRetentionDays"),
			},
			{
				Name:     "default_encryption_scope",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultEncryptionScope"),
			},
			{
				Name:     "deny_encryption_scope_override",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DenyEncryptionScopeOverride"),
			},
			{
				Name:     "public_access",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicAccess"),
			},
			{
				Name:     "last_modified_time",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LastModifiedTime"),
			},
			{
				Name:     "lease_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LeaseStatus"),
			},
			{
				Name:     "lease_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LeaseState"),
			},
			{
				Name:     "lease_duration",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LeaseDuration"),
			},
			{
				Name:     "metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Metadata"),
			},
			{
				Name:     "immutability_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ImmutabilityPolicy"),
			},
			{
				Name:     "legal_hold",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LegalHold"),
			},
			{
				Name:     "has_legal_hold",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HasLegalHold"),
			},
			{
				Name:     "has_immutability_policy",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HasImmutabilityPolicy"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
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

func fetchStorageContainers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Storage.Containers

	account := parent.Item.(storage.Account)
	if !isBlobSupported(&account) {
		return nil
	}

	resource, err := client.ParseResourceID(*account.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	response, err := svc.List(ctx, resource.ResourceGroup, *account.Name, "", "", "")

	if err != nil {
		return errors.WithStack(err)
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
