package dns

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func Searchpaths() *schema.Table {
	return &schema.Table{
		Name:        "tailscale_dns_searchpaths",
		Description: `https://github.com/tailscale/tailscale/blob/main/api.md#tailnet-dns-preferences-get`,
		Resolver:    fetchSearchpaths,
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

type SearchpathResponse struct {
	Name string `json:"name"`
}

func fetchSearchpaths(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	result, err := c.TailscaleClient.DNSSearchPaths(ctx)
	if err != nil {
		return err
	}
	transformedResponse := make([]SearchpathResponse, len(result))
	for i, v := range result {
		transformedResponse[i] = SearchpathResponse{
			Name: v,
		}
	}

	res <- transformedResponse
	return nil
}
