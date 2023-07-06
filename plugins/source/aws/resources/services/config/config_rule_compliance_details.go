package config

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func configRuleComplianceDetails() *schema.Table {
	tableName := "aws_config_config_rule_compliance_details"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/config/latest/APIReference/API_EvaluationResult.html`,
		Resolver:    fetchConfigConfigRuleComplianceDetails,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "config"),
		// no primary key because all the relevant candidate fields can either be null or are not
		// uniquely identifying of a resource. For example, ResourceEvaluationId can be null,
		// and so can ResultToken. However, hashing the entire object can work because a combination of
		// all fields must be unique.
		Transform: transformers.TransformWithStruct(&types.EvaluationResult{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "config_rule_name",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("config_rule_name"),
			},
		},
	}
}

func fetchConfigConfigRuleComplianceDetails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	ruleDetail := parent.Item.(types.ConfigRule)
	cl := meta.(*client.Client)
	svc := cl.Services().Configservice

	input := &configservice.GetComplianceDetailsByConfigRuleInput{
		ConfigRuleName: ruleDetail.ConfigRuleName,
		Limit:          100,
	}
	p := configservice.NewGetComplianceDetailsByConfigRulePaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *configservice.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.EvaluationResults
	}
	return nil
}
