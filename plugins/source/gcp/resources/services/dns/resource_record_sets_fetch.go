package dns

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/dns/v1"
)

func fetchResourceRecordSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	dnsClient, err := dns.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	for {
		managedZone := parent.Item.(*dns.ManagedZone)
		call := dnsClient.ResourceRecordSets.List(c.ProjectId, managedZone.Name)

		if nextPageToken != "" {
			call = call.PageToken(nextPageToken)
		}

		output, err := call.Do()

		if err != nil {
			return err
		}
		res <- output.Rrsets

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
