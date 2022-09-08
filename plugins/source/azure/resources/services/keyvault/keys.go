// Auto generated code - DO NOT EDIT.

package keyvault

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"
)

func keys() *schema.Table {
	return &schema.Table{
		Name:     "azure_keyvault_keys",
		Resolver: fetchKeyVaultKeys,
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
				Name:     "kid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kid"),
			},
			{
				Name:     "attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attributes"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "managed",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Managed"),
			},
		},

		Relations: []*schema.Table{
			keys(),
		},
	}
}

func fetchKeyVaultKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().KeyVault.Keys

	vault := parent.Item.(keyvault.Vault)
	maxResults := int32(25)
	response, err := svc.GetKeys(ctx, *vault.Properties.VaultURI, &maxResults)

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
