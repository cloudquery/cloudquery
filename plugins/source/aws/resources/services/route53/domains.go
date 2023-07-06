package route53

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Domains() *schema.Table {
	tableName := "aws_route53_domains"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetDomainDetail.html`,
		Resolver:            fetchRoute53Domains,
		PreResourceResolver: getDomain,
		Transform:           transformers.TransformWithStruct(&route53domains.GetDomainDetailOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "route53domains"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:       "domain_name",
				Type:       arrow.BinaryTypes.String,
				PrimaryKey: true,
			},
			{
				Name:        "tags",
				Type:        sdkTypes.ExtensionTypes.JSON,
				Resolver:    resolveRoute53DomainTags,
				Description: `A list of tags`,
			},
			{
				Name: "transfer_lock",
				Type: arrow.FixedWidthTypes.Boolean,
			},
		},
	}
}

func fetchRoute53Domains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Route53domains
	var input route53domains.ListDomainsInput
	paginator := route53domains.NewListDomainsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *route53domains.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Domains
	}
	return nil
}
func getDomain(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Route53domains
	v := resource.Item.(types.DomainSummary)

	d, err := svc.GetDomainDetail(ctx, &route53domains.GetDomainDetailInput{DomainName: v.DomainName}, func(options *route53domains.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = d

	return resource.Set("transfer_lock", aws.ToBool(v.TransferLock))
}

func resolveRoute53DomainTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Route53domains
	d := resource.Item.(*route53domains.GetDomainDetailOutput)
	out, err := svc.ListTagsForDomain(ctx, &route53domains.ListTagsForDomainInput{DomainName: d.DomainName}, func(options *route53domains.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(col.Name, client.TagsToMap(out.TagList))
}
