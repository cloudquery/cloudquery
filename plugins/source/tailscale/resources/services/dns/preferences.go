package dns

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func Preferences() *schema.Table {
	return &schema.Table{
		Name:        "tailscale_dns_preferences",
		Description: `https://github.com/tailscale/tailscale/blob/main/api.md#tailnet-dns-preferences-get`,
		Resolver:    fetchPreferences,
		Transform:   transformers.TransformWithStruct(&tailscale.DNSPreferences{}, transformers.WithSkipFields("MagicDNS")),
		Columns: []schema.Column{
			{
				Name:       "tailnet",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveTailnet,
				PrimaryKey: true,
			},
			{
				Name: "magic_dns",
				Type: arrow.FixedWidthTypes.Boolean,
			},
		},
	}
}

func fetchPreferences(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	result, err := c.TailscaleClient.DNSPreferences(ctx)
	if err != nil {
		return err
	}

	res <- result
	return nil
}
