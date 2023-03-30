package elasticsearch

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Packages() *schema.Table {
	tableName := "aws_elasticsearch_packages"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_PackageDetails.html`,
		Resolver:    fetchElasticsearchPackages,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "es"),
		Transform:   client.TransformWithStruct(&types.PackageDetails{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PackageID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
