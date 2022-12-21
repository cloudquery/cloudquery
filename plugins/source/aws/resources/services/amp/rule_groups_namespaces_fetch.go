package amp

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amp"
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

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
