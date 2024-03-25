package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/thoas/go-funk"
)

func resolveRDSTags(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		arn := funk.Get(r.Item, path, funk.WithAllowZero()).(*string)
		cl := meta.(*client.Client)
		svc := cl.Services(client.AWSServiceRds).Rds
		input := rds.ListTagsForResourceInput{ResourceName: arn}
		output, err := svc.ListTagsForResource(ctx, &input, func(options *rds.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		return r.Set(c.Name, client.TagsToMap(output.TagList))
	}
}
