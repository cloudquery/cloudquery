package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func resolveResourceTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, column schema.Column) error {
	c := meta.(*client.Client)
	svc := c.Services().Frauddetector

	paginator := frauddetector.NewListTagsForResourcePaginator(svc,
		&frauddetector.ListTagsForResourceInput{
			ResourceARN: aws.String(resource.Get("arn").String()),
		},
	)
	tags := make(map[string]string)
	for paginator.HasMorePages() {
		data, err := paginator.NextPage(ctx, func(options *frauddetector.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		client.TagsIntoMap(data.Tags, tags)
	}
	return resource.Set(column.Name, tags)
}
