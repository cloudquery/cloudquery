package services

import (
	"context"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/vault/client"
)

func VaultSysPolicies() *schema.Table {
	return &schema.Table{
		Name:     "vault_sys_policies",
		Resolver: fetchVaultSysPolicies,
		Columns: []schema.Column{
			{
				Name:       "policy",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolvePolicy,
				PrimaryKey: true,
			},
		},
	}
}

func fetchVaultSysPolicies(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.VaultServices
	policies, err := svc.Sys.ListPoliciesWithContext(ctx)
	if err != nil {
		return err
	}
	res <- policies
	return nil
}

func resolvePolicy(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return resource.Set(c.Name, resource.Item)
}
