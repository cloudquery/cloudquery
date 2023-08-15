package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func resolveDocDBTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, name, columnName string) error {
	cli := meta.(*client.Client)
	svc := cli.Services().Docdb

	response, err := svc.ListTagsForResource(ctx, &docdb.ListTagsForResourceInput{
		ResourceName: aws.String(name),
	}, func(options *docdb.Options) {
		options.Region = cli.Region
	})
	if err != nil {
		if cli.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(columnName, client.TagsToMap(response.TagList))
}
