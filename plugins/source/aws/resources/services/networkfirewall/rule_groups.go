package networkfirewall

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/networkfirewall/models"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func RuleGroups() *schema.Table {
	tableName := "aws_networkfirewall_rule_groups"
	return &schema.Table{
		Name:                "aws_networkfirewall_rule_groups",
		Description:         `https://docs.aws.amazon.com/network-firewall/latest/APIReference/API_RuleGroup.html`,
		Resolver:            fetchRuleGroups,
		PreResourceResolver: getRuleGroup,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "network-firewall"),
		Transform: transformers.TransformWithStruct(
			&models.RuleGroupWrapper{},
			transformers.WithUnwrapAllEmbeddedStructs(),
		),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RuleGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input networkfirewall.ListRuleGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().Networkfirewall
	p := networkfirewall.NewListRuleGroupsPaginator(svc, &input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *networkfirewall.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}

		res <- response.RuleGroups
	}
	return nil
}

func getRuleGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Networkfirewall
	metadata := resource.Item.(types.RuleGroupMetadata)

	ruleGroup, err := svc.DescribeRuleGroup(ctx, &networkfirewall.DescribeRuleGroupInput{
		RuleGroupArn: metadata.Arn,
	}, func(options *networkfirewall.Options) {
		options.Region = c.Region
	})
	if err != nil && !c.IsNotFoundError(err) {
		return err
	}

	resource.Item = &models.RuleGroupWrapper{
		RuleGroup:         ruleGroup.RuleGroup,
		RuleGroupResponse: ruleGroup.RuleGroupResponse,
	}
	return nil
}
