package elasticsearch

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func Versions() *schema.Table {
	tableName := "aws_elasticsearch_versions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_ListVersions.html`,
		Resolver:    fetchElasticsearchVersions,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "es"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "version",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveVersion,
				PrimaryKey: true,
			},
			{
				Name:     "instance_types",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveInstanceTypes,
			},
		},
	}
}

func fetchElasticsearchVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Elasticsearchservice

	p := elasticsearchservice.NewListElasticsearchVersionsPaginator(svc,
		&elasticsearchservice.ListElasticsearchVersionsInput{MaxResults: 100},
	)
	for p.HasMorePages() {
		out, err := p.NextPage(ctx, func(options *elasticsearchservice.Options) {
			options.Region = cl.Region
		})
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
	cl := meta.(*client.Client)
	svc := cl.Services().Elasticsearchservice

	var instanceTypes []types.ESPartitionInstanceType
	p := elasticsearchservice.NewListElasticsearchInstanceTypesPaginator(svc,
		&elasticsearchservice.ListElasticsearchInstanceTypesInput{
			ElasticsearchVersion: aws.String(resource.Item.(string)),
		},
	)
	for p.HasMorePages() {
		out, err := p.NextPage(ctx, func(options *elasticsearchservice.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		instanceTypes = append(instanceTypes, out.ElasticsearchInstanceTypes...)
	}

	return resource.Set(c.Name, instanceTypes)
}
