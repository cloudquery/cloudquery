package elasticsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Domains() *schema.Table {
	tableName := "aws_elasticsearch_domains"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_DomainStatus.html`,
		Resolver:            fetchElasticsearchDomains,
		PreResourceResolver: getDomain,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "es"),
		Transform:           transformers.TransformWithStruct(&types.ElasticsearchDomainStatus{}),
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

func fetchElasticsearchDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Elasticsearchservice
	out, err := svc.ListDomainNames(ctx, nil)
	if err != nil {
		return err
	}

	res <- out.DomainNames

	return nil
}

func getDomain(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	svc := meta.(*client.Client).Services().Elasticsearchservice

	out, err := svc.DescribeElasticsearchDomain(ctx,
		&elasticsearchservice.DescribeElasticsearchDomainInput{
			DomainName: resource.Item.(types.DomainInfo).DomainName,
		},
	)
	if err != nil {
		return err
	}

	resource.SetItem(out.DomainStatus)

	return nil
}

func resolveDomainTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	svc := meta.(*client.Client).Services().Elasticsearchservice

	tagsOutput, err := svc.ListTags(ctx,
		&elasticsearchservice.ListTagsInput{
			ARN: resource.Item.(*types.ElasticsearchDomainStatus).ARN,
		},
	)
	if err != nil {
		return err
	}

	return resource.Set(c.Name, client.TagsToMap(tagsOutput.TagList))
}

func resolveAuthorizedPrincipals(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	svc := meta.(*client.Client).Services().Elasticsearchservice

	input := &elasticsearchservice.ListVpcEndpointAccessInput{
		DomainName: resource.Item.(*types.ElasticsearchDomainStatus).DomainName,
	}

	var principals []types.AuthorizedPrincipal
	// No paginator available
	for {
		out, err := svc.ListVpcEndpointAccess(ctx, input)
		if err != nil {
			return err
		}

		principals = append(principals, out.AuthorizedPrincipalList...)

		if aws.ToString(out.NextToken) == "" {
			break
		}

		input.NextToken = out.NextToken
	}

	return resource.Set(c.Name, principals)
}
