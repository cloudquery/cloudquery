package elasticsearch

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func VpcEndpoints() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticsearch_vpc_endpoints",
		Description: `https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_VpcEndpoint.html`,
		Resolver:    fetchElasticsearchVpcEndpoints,
		Multiplex:   client.ServiceAccountRegionMultiplexer("es"),
		Transform:   transformers.TransformWithStruct(&types.VpcEndpoint{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcEndpointId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
