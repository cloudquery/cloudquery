package services

import (
	"context"
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/vault/client"
	"github.com/hashicorp/vault/api"
)

type sysMountWithPath struct {
	path string
	*api.MountOutput
}

func VaultSysMounts() *schema.Table {
	return &schema.Table{
		Name:                "vault_sys_mounts",
		Description:         "https://developer.hashicorp.com/vault/api-docs/system/mounts",
		Resolver:            fetchVaultSysMounts,
		PreResourceResolver: resolveVaultSysMount,
		Transform:           transformers.TransformWithStruct(api.MountOutput{}, transformers.WithPrimaryKeys("UUID")),
		Columns: []schema.Column{
			{
				Name:        "path",
				Type:        arrow.BinaryTypes.String,
				Description: "The mount path for the sys mount resource",
			},
		},
	}
}

func fetchVaultSysMounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.VaultServices
	mounts, err := svc.Sys.ListMountsWithContext(ctx)
	if err != nil {
		return err
	}
	for path, mount := range mounts {
		res <- sysMountWithPath{
			path:        path,
			MountOutput: mount,
		}
	}
	return nil
}

func resolveVaultSysMount(_ context.Context, _ schema.ClientMeta, resource *schema.Resource) error {
	mountWithPath := resource.Item.(sysMountWithPath)
	resource.Item = mountWithPath.MountOutput
	return resource.Set("path", mountWithPath.path)
}
