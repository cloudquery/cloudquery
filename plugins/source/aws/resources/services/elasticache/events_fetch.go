package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchElasticacheEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input elasticache.DescribeEventsInput

	paginator := elasticache.NewDescribeEventsPaginator(meta.(*client.Client).Services().Elasticache, &input)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- v.Events
	}
	return nil
}
