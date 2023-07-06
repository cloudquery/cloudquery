package route53resolver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func FirewallDomainLists() *schema.Table {
	tableName := "aws_route53resolver_firewall_domain_lists"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_FirewallDomainList.html`,
		Resolver:            fetchFirewallDomainList,
		PreResourceResolver: getFirewallDomainList,
		Transform:           transformers.TransformWithStruct(&types.FirewallDomainList{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "route53resolver"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}

func fetchFirewallDomainList(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Route53resolver
	var input route53resolver.ListFirewallDomainListsInput
	paginator := route53resolver.NewListFirewallDomainListsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *route53resolver.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.FirewallDomainLists
	}
	return nil
}
func getFirewallDomainList(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Route53resolver
	v := resource.Item.(types.FirewallDomainListMetadata)

	d, err := svc.GetFirewallDomainList(ctx, &route53resolver.GetFirewallDomainListInput{FirewallDomainListId: v.Id}, func(options *route53resolver.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = d.FirewallDomainList

	return nil
}
