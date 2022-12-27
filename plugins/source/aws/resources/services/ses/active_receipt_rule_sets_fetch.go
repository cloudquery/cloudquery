package ses

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

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

	if !isRegionSupported(c.Region) {
		return nil
	}

	set, err := svc.DescribeActiveReceiptRuleSet(ctx, nil)
	if err != nil {
		return err
	}
	res <- set

	return nil
}
