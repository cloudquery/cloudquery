package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchConfigConfigRuleCompliances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	ruleDetail := parent.Item.(types.ConfigRule)
	c := meta.(*client.Client)
	svc := c.Services().Configservice

	input := &configservice.DescribeComplianceByConfigRuleInput{
		ConfigRuleNames: []string{aws.ToString(ruleDetail.ConfigRuleName)},
	}
	p := configservice.NewDescribeComplianceByConfigRulePaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.ComplianceByConfigRules
	}
	return nil
}
