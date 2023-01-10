package elasticsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

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
	for {
		out, err := svc.ListVpcEndpointAccess(ctx, input)
		if err != nil {
			return err
		}

		principals = append(principals, out.AuthorizedPrincipalList...)

		if out.NextToken == nil {
			break
		}

		input.NextToken = out.NextToken
	}

	return resource.Set(c.Name, principals)
}
