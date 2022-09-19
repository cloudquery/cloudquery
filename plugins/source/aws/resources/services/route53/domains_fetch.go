package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRoute53Domains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Route53Domains
	var input route53domains.ListDomainsInput
	optsFunc := func(options *route53domains.Options) {
		// Set region to default global region
		options.Region = "us-east-1"
	}
	for {
		output, err := svc.ListDomains(ctx, &input, optsFunc)
		if err != nil {
			return err
		}

		for _, v := range output.Domains {
			d, err := svc.GetDomainDetail(ctx, &route53domains.GetDomainDetailInput{DomainName: v.DomainName}, optsFunc)
			if err != nil {
				return err
			}
			res <- d
		}

		if aws.ToString(output.NextPageMarker) == "" {
			break
		}
		input.Marker = output.NextPageMarker
	}
	return nil
}

func resolveRoute53DomainTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	c := meta.(*client.Client)
	svc := c.Services().Route53Domains
	d := resource.Item.(*route53domains.GetDomainDetailOutput)
	out, err := svc.ListTagsForDomain(ctx, &route53domains.ListTagsForDomainInput{DomainName: d.DomainName}, func(options *route53domains.Options) {
		// Set region to default global region
		options.Region = "us-east-1"
	})
	if err != nil {
		return err
	}
	return resource.Set(col.Name, client.TagsToMap(out.TagList))
}
