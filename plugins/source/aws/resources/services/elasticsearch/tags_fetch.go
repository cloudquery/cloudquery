package elasticsearch

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func resolveTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().Elasticsearchservice

	var arn *string
	switch typed := resource.Item.(type) {
	case *types.ElasticsearchDomainStatus:
		arn = typed.ARN
	default:
		return fmt.Errorf("unsupported type for resolveTags: %T", typed)
	}

	tagsOutput, err := svc.ListTags(ctx, &elasticsearchservice.ListTagsInput{
		ARN: arn,
	}, func(o *elasticsearchservice.Options) {
		o.Region = region
	})
	if err != nil {
		return err
	}

	return resource.Set(c.Name, client.TagsToMap(tagsOutput.TagList))
}
