package ses

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSesActiveReceiptRuleSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Ses

	set, err := svc.DescribeActiveReceiptRuleSet(ctx, nil)
	if err != nil {
		return err
	}
	res <- set

	return nil
}
