package route53

import (
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Domains() *schema.Table {
	return &schema.Table{
		Name:                "aws_route53_domains",
		Description:         `https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetDomainDetail.html`,
		Resolver:            fetchRoute53Domains,
		PreResourceResolver: getDomain,
		Transform:           transformers.TransformWithStruct(&route53domains.GetDomainDetailOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("route53domains"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name: "domain_name",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveRoute53DomainTags,
				Description: `A list of tags`,
			},
			{
				Name: "transfer_lock",
				Type: schema.TypeBool,
			},
		},
	}
}
