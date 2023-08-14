package services

import (
	"context"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/vault/client"
	"github.com/hashicorp/vault/api"
)

func VaultSysPlugins() *schema.Table {
	return &schema.Table{
		Name:      "vault_sys_plugins",
		Resolver:  fetchVaultSysPlugins,
		Transform: transformers.TransformWithStruct(api.PluginDetails{}, transformers.WithPrimaryKeys("Type", "Name", "Version")),
	}
}

func fetchVaultSysPlugins(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.VaultServices
	response, err := svc.Sys.ListPluginsWithContext(ctx, &api.ListPluginsInput{})
	if err != nil {
		return err
	}
	res <- response.Details
	return nil
}
