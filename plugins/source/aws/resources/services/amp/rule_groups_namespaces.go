package amp

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amp"
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Arn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchAmpRuleGroupsNamespaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Amp

	p := amp.NewListRuleGroupsNamespacesPaginator(svc,
		&amp.ListRuleGroupsNamespacesInput{
			WorkspaceId: parent.Item.(*types.WorkspaceDescription).WorkspaceId,
			MaxResults:  aws.Int32(int32(1000)),
		},
	)
	for p.HasMorePages() {
		out, err := p.NextPage(ctx,
			func(options *amp.Options) {
				options.Region = cl.Region
			})
		if err != nil {
			return err
		}

		res <- out.RuleGroupsNamespaces
	}

	return nil
}

func describeRuleGroupsNamespace(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Amp

	out, err := svc.DescribeRuleGroupsNamespace(ctx,
		&amp.DescribeRuleGroupsNamespaceInput{WorkspaceId: resource.Parent.Item.(*types.WorkspaceDescription).WorkspaceId},
		func(options *amp.Options) {
			options.Region = cl.Region
		},
	)
	if err != nil {
		return err
	}

	resource.SetItem(out.RuleGroupsNamespace)

	return nil
}
