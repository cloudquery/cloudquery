package acl

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/cloudquery/plugin-sdk/v3/types"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func Acls() *schema.Table {
	return &schema.Table{
		Name:        "tailscale_acls",
		Description: `https://github.com/tailscale/tailscale/blob/main/api.md#acl`,
		Resolver:    fetchAcls,
		Transform:   transformers.TransformWithStruct(&tailscale.ACL{}, transformers.WithSkipFields("ACLs", "DERPMap", "SSH", "DisableIPv4", "OneCGNATRoute")),
		Columns: []schema.Column{
			{
				Name:       "tailnet",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveTailnet,
				PrimaryKey: true,
			},
			{
				Name:     "acls",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("ACLs"),
			},
			{
				Name:     "derp_map",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("DERPMap"),
			},
			{
				Name:     "ssh",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("SSH"),
			},
			{
				Name:     "disable_ipv4",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: schema.PathResolver("DisableIPv4"),
			},
			{
				Name:     "one_cgnat_route",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("OneCGNATRoute"),
			},
		},
	}
}

func fetchAcls(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	result, err := c.TailscaleClient.ACL(ctx)
	if err != nil {
		return err
	}

	res <- result
	return nil
}
