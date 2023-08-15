package quicksight

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/thoas/go-funk"
)

var tagsCol = schema.Column{
	Name:     "tags",
	Type:     sdkTypes.ExtensionTypes.JSON,
	Resolver: resolveTags,
}

func resolveTags(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	arn := funk.Get(r.Item, "Arn", funk.WithAllowZero()).(*string)
	cl := meta.(*client.Client)
	svc := cl.Services().Quicksight
	params := quicksight.ListTagsForResourceInput{ResourceArn: arn}

	output, err := svc.ListTagsForResource(ctx, &params, func(options *quicksight.Options) {
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
