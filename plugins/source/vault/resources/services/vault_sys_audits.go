package services

import (
	"context"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/vault/client"
	"github.com/hashicorp/vault/api"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func VaultSysAudits() *schema.Table {
	return &schema.Table{
		Name:        "vault_sys_audits",
		Description: "https://developer.hashicorp.com/vault/api-docs/system/audit",
		Resolver:    fetchVaultSysAudits,
		Transform:   transformers.TransformWithStruct(api.Audit{}, transformers.WithPrimaryKeys("Type", "Path")),
	}
}

func fetchVaultSysAudits(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.VaultServices
	audits, err := svc.Sys.ListAuditWithContext(ctx)
	if err != nil {
		return err
	}
	for _, audit := range audits {
		res <- audit
	}
	return nil
}
