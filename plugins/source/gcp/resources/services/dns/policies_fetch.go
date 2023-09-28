package dns

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"google.golang.org/api/dns/v1"
)

func fetchPolicies(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	dnsClient, err := dns.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	for {
		output, err := dnsClient.Policies.List(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		res <- output.Policies

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
