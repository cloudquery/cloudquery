package elasticsearch

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Packages() *schema.Table {
	tableName := "aws_elasticsearch_packages"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_PackageDetails.html`,
		Resolver:    fetchElasticsearchPackages,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "es"),
		Transform:   transformers.TransformWithStruct(&types.PackageDetails{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("PackageID"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchElasticsearchPackages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Elasticsearchservice

	p := elasticsearchservice.NewDescribePackagesPaginator(svc, nil)
	for p.HasMorePages() {
		out, err := p.NextPage(ctx, func(options *elasticsearchservice.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- out.PackageDetailsList
	}

	return nil
}
