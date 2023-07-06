package networkfirewall

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/networkfirewall/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func RuleGroups() *schema.Table {
	tableName := "aws_networkfirewall_rule_groups"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/network-firewall/latest/APIReference/API_RuleGroup.html`,
		Resolver:            fetchRuleGroups,
		PreResourceResolver: getRuleGroup,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "network-firewall"),
		Transform: transformers.TransformWithStruct(
			&models.RuleGroupWrapper{},
			transformers.WithUnwrapAllEmbeddedStructs(),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("RuleGroupArn"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input networkfirewall.ListRuleGroupsInput
	cl := meta.(*client.Client)
	svc := cl.Services().Networkfirewall
	p := networkfirewall.NewListRuleGroupsPaginator(svc, &input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *networkfirewall.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- response.RuleGroups
	}
	return nil
}

func getRuleGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Networkfirewall
	metadata := resource.Item.(types.RuleGroupMetadata)

	ruleGroup, err := svc.DescribeRuleGroup(ctx, &networkfirewall.DescribeRuleGroupInput{
		RuleGroupArn: metadata.Arn,
	}, func(options *networkfirewall.Options) {
		options.Region = cl.Region
	})
	if err != nil && !cl.IsNotFoundError(err) {
		return err
	}

	resource.Item = &models.RuleGroupWrapper{
		RuleGroup:         ruleGroup.RuleGroup,
		RuleGroupResponse: ruleGroup.RuleGroupResponse,
	}
	return nil
}
