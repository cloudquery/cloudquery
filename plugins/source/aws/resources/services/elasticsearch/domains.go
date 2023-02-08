package elasticsearch

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Domains() *schema.Table {
	return &schema.Table{
		Name:                "aws_elasticsearch_domains",
		Description:         `https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_DomainStatus.html`,
		Resolver:            fetchElasticsearchDomains,
		PreResourceResolver: getDomain,
		Multiplex:           client.ServiceAccountRegionMultiplexer("es"),
		Transform:           transformers.TransformWithStruct(&types.ElasticsearchDomainStatus{}),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				RetainOrder: true,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
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
