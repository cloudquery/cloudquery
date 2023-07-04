package dns

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func Nameservers() *schema.Table {
	return &schema.Table{
		Name:        "tailscale_dns_nameservers",
		Description: `https://github.com/tailscale/tailscale/blob/main/api.md#tailnet-dns-preferences-get`,
		Resolver:    fetchNameservers,
		Columns: []schema.Column{
			{
				Name:       "tailnet",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveTailnet,
				PrimaryKey: true,
			},
			{
				Name:       "name",
				Type:       arrow.BinaryTypes.String,
				PrimaryKey: true,
			},
		},
	}
}

type NameserverResponse struct {
	Name string `json:"name"`
}

func fetchNameservers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	result, err := c.TailscaleClient.DNSNameservers(ctx)
	if err != nil {
		return err
	}
	transformedResponse := make([]NameserverResponse, len(result))
	for i, v := range result {
		transformedResponse[i] = NameserverResponse{
			Name: v,
		}
	}
	res <- transformedResponse
	return nil
}
