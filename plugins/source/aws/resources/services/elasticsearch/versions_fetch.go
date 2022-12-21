package elasticsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchElasticsearchVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Elasticsearchservice

	p := elasticsearchservice.NewListElasticsearchVersionsPaginator(svc,
		&elasticsearchservice.ListElasticsearchVersionsInput{MaxResults: 100},
	)
	for p.HasMorePages() {
		out, err := p.NextPage(ctx)
		if err != nil {
			return err
		}

		res <- out.ElasticsearchVersions
	}

	return nil
}

func resolveVersion(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return resource.Set(c.Name, resource.Item.(string))
}

func resolveInstanceTypes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	svc := meta.(*client.Client).Services().Elasticsearchservice

	var instanceTypes []types.ESPartitionInstanceType
	p := elasticsearchservice.NewListElasticsearchInstanceTypesPaginator(svc,
		&elasticsearchservice.ListElasticsearchInstanceTypesInput{
			ElasticsearchVersion: aws.String(resource.Item.(string)),
		},
	)
	for p.HasMorePages() {
		out, err := p.NextPage(ctx)
		if err != nil {
			return err
		}

		instanceTypes = append(instanceTypes, out.ElasticsearchInstanceTypes...)
	}

	return resource.Set(c.Name, instanceTypes)
}
