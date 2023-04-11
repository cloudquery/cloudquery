package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/services"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func resolveIotTags(ctx context.Context, svc services.IotClient, resource *schema.Resource, c schema.Column, resourceArn *string) error {
	input := iot.ListTagsForResourceInput{
		ResourceArn: resourceArn,
	}
	tags := make(map[string]string)
	paginator := iot.NewListTagsForResourcePaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		client.TagsIntoMap(page.Tags, tags)
	}
	return resource.Set(c.Name, tags)
}
