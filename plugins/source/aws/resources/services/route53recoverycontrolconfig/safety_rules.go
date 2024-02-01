package route53recoverycontrolconfig

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func safetyRules() *schema.Table {
	tableName := "aws_route53recoverycontrolconfig_safety_rules"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/recovery-cluster/latest/api/controlpanel-controlpanelarn-safetyrules.html`,
		Resolver:    fetchSafetyRules,
		Transform:   transformers.TransformWithStruct(&types.Rule{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveRuleARN,
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchSafetyRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRoute53recoverycontrolconfig).Route53recoverycontrolconfig
	controlPanel := parent.Item.(types.ControlPanel)
	input := &route53recoverycontrolconfig.ListSafetyRulesInput{
		ControlPanelArn: controlPanel.ControlPanelArn,
	}
	paginator := route53recoverycontrolconfig.NewListSafetyRulesPaginator(svc, input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *route53recoverycontrolconfig.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.SafetyRules
	}
	return nil
}

func resolveRuleARN(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	rule := r.Item.(types.Rule)
	// A safety rule can be an assertion rule or a gating rule so only one of the following will be non-nil at any time
	if rule.ASSERTION != nil {
		return r.Set(c.Name, rule.ASSERTION.SafetyRuleArn)
	}
	if rule.GATING != nil {
		return r.Set(c.Name, rule.GATING.SafetyRuleArn)
	}
	return nil
}
