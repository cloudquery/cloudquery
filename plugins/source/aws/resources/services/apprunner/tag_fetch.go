package apprunner

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/thoas/go-funk"
)

func resolveApprunnerTags(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		arn := funk.Get(r.Item, path, funk.WithAllowZero()).(*string)
		cl := meta.(*client.Client)
		svc := cl.Services().Apprunner
		params := apprunner.ListTagsForResourceInput{ResourceArn: arn}

		output, err := svc.ListTagsForResource(ctx, &params)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		return r.Set(c.Name, client.TagsToMap(output.Tags))
	}
}
