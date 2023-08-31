package services

import (
	"context"
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/vault/client"
	"github.com/hashicorp/vault/api"
)

type authMountWithPath struct {
	path string
	*api.AuthMount
}

func VaultSysAuths() *schema.Table {
	return &schema.Table{
		Name:                "vault_sys_auths",
		Description:         "https://developer.hashicorp.com/vault/api-docs/system/auth",
		Resolver:            fetchVaultSysAuths,
		PreResourceResolver: resolveVaultSysAuth,
		Transform:           transformers.TransformWithStruct(api.AuthMount{}, transformers.WithPrimaryKeys("UUID")),
		Columns: []schema.Column{
			{
				Name:        "path",
				Type:        arrow.BinaryTypes.String,
				Description: "The mount path for the auth resource",
			},
		},
	}
}

func fetchVaultSysAuths(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.VaultServices
	auths, err := svc.Sys.ListAuthWithContext(ctx)
	if err != nil {
		return err
	}
	for path, auth := range auths {
		res <- authMountWithPath{
			path:      path,
			AuthMount: auth,
		}
	}
	return nil
}

func resolveVaultSysAuth(_ context.Context, _ schema.ClientMeta, resource *schema.Resource) error {
	mountWithPath := resource.Item.(authMountWithPath)
	resource.Item = mountWithPath.AuthMount
	return resource.Set("path", mountWithPath.path)
}
