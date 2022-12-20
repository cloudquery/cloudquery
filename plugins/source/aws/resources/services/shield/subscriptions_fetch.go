package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchShieldSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Shield
	config := shield.DescribeSubscriptionInput{}
	output, err := svc.DescribeSubscription(ctx, &config)
	if err != nil {
		if c.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	res <- output.Subscription
	return nil
}
