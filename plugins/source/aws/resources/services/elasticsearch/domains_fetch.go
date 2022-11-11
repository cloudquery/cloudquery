package elasticsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchElasticsearchDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Elasticsearchservice
	out, err := svc.ListDomainNames(ctx, &elasticsearchservice.ListDomainNamesInput{})
	if err != nil {
		return err
	}

	res <- out.DomainNames
	return nil
}

func getDomain(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Elasticsearchservice

	info := resource.Item.(types.DomainInfo)

	domainOutput, err := svc.DescribeElasticsearchDomain(ctx, &elasticsearchservice.DescribeElasticsearchDomainInput{DomainName: info.DomainName})
	if err != nil {
		return nil
	}

	resource.Item = domainOutput.DomainStatus
	return nil
}

func resolveElasticsearchDomainTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().Elasticsearchservice
	domain := resource.Item.(*types.ElasticsearchDomainStatus)
	tagsOutput, err := svc.ListTags(ctx, &elasticsearchservice.ListTagsInput{
		ARN: domain.ARN,
	}, func(o *elasticsearchservice.Options) {
		o.Region = region
	})
	if err != nil {
		return err
	}
	if len(tagsOutput.TagList) == 0 {
		return nil
	}
	tags := make(map[string]*string)
	for _, s := range tagsOutput.TagList {
		tags[*s.Key] = s.Value
	}
	return resource.Set(c.Name, tags)
}
