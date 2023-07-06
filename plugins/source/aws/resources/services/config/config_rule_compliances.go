package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func configRuleCompliances() *schema.Table {
	tableName := "aws_config_config_rule_compliances"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/config/latest/APIReference/API_ComplianceByConfigRule.html`,
		Resolver:    fetchConfigConfigRuleCompliances,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "config"),
		Transform:   transformers.TransformWithStruct(&types.ComplianceByConfigRule{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}

func fetchConfigConfigRuleCompliances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	ruleDetail := parent.Item.(types.ConfigRule)
	cl := meta.(*client.Client)
	svc := cl.Services().Configservice

	input := &configservice.DescribeComplianceByConfigRuleInput{
		ConfigRuleNames: []string{aws.ToString(ruleDetail.ConfigRuleName)},
	}
	p := configservice.NewDescribeComplianceByConfigRulePaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *configservice.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.ComplianceByConfigRules
	}
	return nil
}
