package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func domainClientOpts(options *route53domains.Options) {
	// Set region to default global region
	options.Region = "us-east-1"
}

func fetchRoute53Domains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Route53domains
	var input route53domains.ListDomainsInput

	for {
		output, err := svc.ListDomains(ctx, &input, domainClientOpts)
		if err != nil {
			return err
		}
		res <- output.Domains

		if aws.ToString(output.NextPageMarker) == "" {
			break
		}
		input.Marker = output.NextPageMarker
	}
	return nil
}
func getDomain(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Route53domains
	v := resource.Item.(types.DomainSummary)

	d, err := svc.GetDomainDetail(ctx, &route53domains.GetDomainDetailInput{DomainName: v.DomainName}, domainClientOpts)
	if err != nil {
		return err
	}

	resource.Item = d

	return resource.Set("transfer_lock", aws.ToBool(v.TransferLock))
}

func resolveRoute53DomainTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	c := meta.(*client.Client)
	svc := c.Services().Route53domains
	d := resource.Item.(*route53domains.GetDomainDetailOutput)
	out, err := svc.ListTagsForDomain(ctx, &route53domains.ListTagsForDomainInput{DomainName: d.DomainName}, domainClientOpts)
	if err != nil {
		return err
	}
	return resource.Set(col.Name, client.TagsToMap(out.TagList))
}
