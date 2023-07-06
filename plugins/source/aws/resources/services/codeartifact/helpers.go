package codeartifact

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"

	"github.com/thoas/go-funk"
)

func resolveCodeartifactTags(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		arn := funk.Get(r.Item, path, funk.WithAllowZero()).(*string)
		cl := meta.(*client.Client)
		svc := cl.Services().Codeartifact
		params := codeartifact.ListTagsForResourceInput{ResourceArn: arn}

		output, err := svc.ListTagsForResource(ctx, &params, func(options *codeartifact.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		return r.Set(c.Name, client.TagsToMap(output.Tags))
	}
}
