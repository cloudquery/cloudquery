package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/thoas/go-funk"
)

var tagsCol = schema.Column{
	Name:     "tags",
	Type:     schema.TypeJSON,
	Resolver: resolveTags,
}

func resolveTags(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	arn := funk.Get(r.Item, "Arn", funk.WithAllowZero()).(*string)
	cl := meta.(*client.Client)
	svc := cl.Services().Quicksight
	params := quicksight.ListTagsForResourceInput{ResourceArn: arn}

	output, err := svc.ListTagsForResource(ctx, &params)
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return r.Set(c.Name, client.TagsToMap(output.Tags))
}
