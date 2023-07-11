package route53resolver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func FirewallRuleGroupAssociations() *schema.Table {
	tableName := "aws_route53resolver_firewall_rule_group_associations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_FirewallRuleGroupAssociation.html`,
		Resolver:    fetchFirewallRuleGroupAssociations,
		Transform:   transformers.TransformWithStruct(&types.FirewallRuleGroupAssociation{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "route53resolver"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}

func fetchFirewallRuleGroupAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Route53resolver
	var input route53resolver.ListFirewallRuleGroupAssociationsInput
	paginator := route53resolver.NewListFirewallRuleGroupAssociationsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *route53resolver.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.FirewallRuleGroupAssociations
	}
	return nil
}
