package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func resolveResourceTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, column schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Frauddetector

	paginator := frauddetector.NewListTagsForResourcePaginator(svc,
		&frauddetector.ListTagsForResourceInput{
			ResourceARN: aws.String(resource.Get("arn").String()),
		},
	)
	tags := make(map[string]string)
	for paginator.HasMorePages() {
		data, err := paginator.NextPage(ctx, func(options *frauddetector.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		client.TagsIntoMap(data.Tags, tags)
	}
	return resource.Set(column.Name, tags)
}
