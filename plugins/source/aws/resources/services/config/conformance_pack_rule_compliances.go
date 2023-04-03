package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/config/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func conformancePackRuleCompliances() *schema.Table {
	tableName := "aws_config_conformance_pack_rule_compliances"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/config/latest/APIReference/API_DescribeConformancePackCompliance.html`,
		Resolver:    fetchConfigConformancePackRuleCompliances,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "config"),
		Transform:   transformers.TransformWithStruct(&models.ConformancePackComplianceWrapper{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "conformance_pack_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchConfigConformancePackRuleCompliances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	conformancePackDetail := parent.Item.(types.ConformancePackDetail)
	c := meta.(*client.Client)
	cs := c.Services().Configservice
	params := configservice.DescribeConformancePackComplianceInput{
		ConformancePackName: conformancePackDetail.ConformancePackName,
	}
	for {
		resp, err := cs.DescribeConformancePackCompliance(ctx, &params)
		if err != nil {
			return err
		}
		for _, conformancePackRuleCompliance := range resp.ConformancePackRuleComplianceList {
			detailParams := &configservice.GetConformancePackComplianceDetailsInput{
				ConformancePackName: conformancePackDetail.ConformancePackName,
				Filters: &types.ConformancePackEvaluationFilters{
					ConfigRuleNames: []string{*conformancePackRuleCompliance.ConfigRuleName},
				},
			}
			for {
				output, err := cs.GetConformancePackComplianceDetails(ctx, detailParams)
				if err != nil {
					return err
				}
				for _, conformancePackComplianceDetail := range output.ConformancePackRuleEvaluationResults {
					res <- models.ConformancePackComplianceWrapper{
						ComplianceType:             conformancePackRuleCompliance.ComplianceType,
						ConfigRuleName:             conformancePackRuleCompliance.ConfigRuleName,
						Controls:                   conformancePackRuleCompliance.Controls,
						ConfigRuleInvokedTime:      conformancePackComplianceDetail.ConfigRuleInvokedTime,
						EvaluationResultIdentifier: conformancePackComplianceDetail.EvaluationResultIdentifier,
						ResultRecordedTime:         conformancePackComplianceDetail.ResultRecordedTime,
						Annotation:                 conformancePackComplianceDetail.Annotation,
					}
				}
				if output.NextToken == nil {
					break
				}
				detailParams.NextToken = output.NextToken
			}
		}
		if resp.NextToken == nil {
			break
		}
		params.NextToken = resp.NextToken
	}
	return nil
}
