package amp

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amp"
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func ruleGroupsNamespaces() *schema.Table {
	return &schema.Table{
		Name:                "aws_amp_rule_groups_namespaces",
		Description:         `https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-APIReference.html#AMP-APIReference-RuleGroupsNamespaceDescription`,
		Resolver:            fetchAmpRuleGroupsNamespaces,
		PreResourceResolver: describeRuleGroupsNamespace,
		Transform:           transformers.TransformWithStruct(&types.RuleGroupsNamespaceDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "workspace_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchAmpRuleGroupsNamespaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Amp

	p := amp.NewListRuleGroupsNamespacesPaginator(svc,
		&amp.ListRuleGroupsNamespacesInput{
			WorkspaceId: parent.Item.(*types.WorkspaceDescription).WorkspaceId,
			MaxResults:  aws.Int32(int32(1000)),
		},
	)
	for p.HasMorePages() {
		out, err := p.NextPage(ctx)
		if err != nil {
			return err
		}

		res <- out.RuleGroupsNamespaces
	}

	return nil
}

func describeRuleGroupsNamespace(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	svc := meta.(*client.Client).Services().Amp

	out, err := svc.DescribeRuleGroupsNamespace(ctx,
		&amp.DescribeRuleGroupsNamespaceInput{WorkspaceId: resource.Parent.Item.(*types.WorkspaceDescription).WorkspaceId},
	)
	if err != nil {
		return err
	}

	resource.SetItem(out.RuleGroupsNamespace)

	return nil
}
