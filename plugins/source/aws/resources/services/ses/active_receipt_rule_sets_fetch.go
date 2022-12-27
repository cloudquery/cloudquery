package ses

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

// Supported regions based on https://docs.aws.amazon.com/ses/latest/dg/regions.html#region-receive-email
// We hard code as there isn't a good way to automatically fetch this list
var supportedRegions = []string{"us-east-1", "us-west-2", "eu-west-1"}

func isRegionSupported(region string) bool {
	for _, r := range supportedRegions {
		if r == region {
			return true
		}
	}
	return false
}

func fetchSesActiveReceiptRuleSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Ses

	set, err := svc.DescribeActiveReceiptRuleSet(ctx, nil)
	if err != nil {
		if !isRegionSupported(c.Region) && client.IgnoreWithInvalidAction(err) {
			return nil
		}
		return err
	}

	if set.Metadata != nil && set.Metadata.Name != nil {
		res <- set
	}

	return nil
}
