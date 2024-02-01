package config

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
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
		Transform:   transformers.TransformWithStruct(&types.ComplianceByConfigRule{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "config_rule_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchConfigConfigRuleCompliances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	ruleDetail := parent.Item.(types.ConfigRule)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceConfigservice).Configservice

	input := &configservice.DescribeComplianceByConfigRuleInput{
		ConfigRuleNames: []string{aws.ToString(ruleDetail.ConfigRuleName)}, // so we'll have only a single result
	}

	// we request a single config rule info, so no need to iterate
	response, err := svc.DescribeComplianceByConfigRule(ctx, input, func(options *configservice.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- response.ComplianceByConfigRules
	return nil
}
