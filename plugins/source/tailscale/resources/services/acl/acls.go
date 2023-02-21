package acl

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Name:     "tailnet",
				Type:     schema.TypeString,
				Resolver: client.ResolveTailnet,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "acls",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ACLs"),
			},
			{
				Name:     "derp_map",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DERPMap"),
			},
			{
				Name:     "ssh",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SSH"),
			},
			{
				Name:     "disable_ipv4",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DisableIPv4"),
			},
			{
				Name:     "one_cgnat_route",
				Type:     schema.TypeString,
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
