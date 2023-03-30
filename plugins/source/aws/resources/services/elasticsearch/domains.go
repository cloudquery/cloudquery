package elasticsearch

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Domains() *schema.Table {
	tableName := "aws_elasticsearch_domains"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_DomainStatus.html`,
		Resolver:            fetchElasticsearchDomains,
		PreResourceResolver: getDomain,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "es"),
		Transform:           client.TransformWithStruct(&types.ElasticsearchDomainStatus{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "authorized_principals",
				Type:     schema.TypeJSON,
				Resolver: resolveAuthorizedPrincipals,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveDomainTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
