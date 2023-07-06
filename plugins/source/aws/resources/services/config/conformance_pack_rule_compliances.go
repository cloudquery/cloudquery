package config

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/config/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchConfigConformancePackRuleCompliances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	conformancePackDetail := parent.Item.(types.ConformancePackDetail)
	cl := meta.(*client.Client)
	cs := cl.Services().Configservice
	params := configservice.DescribeConformancePackComplianceInput{
		ConformancePackName: conformancePackDetail.ConformancePackName,
	}
	paginator := configservice.NewDescribeConformancePackCompliancePaginator(cs, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *configservice.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		for _, conformancePackRuleCompliance := range page.ConformancePackRuleComplianceList {
			detailParams := &configservice.GetConformancePackComplianceDetailsInput{
				ConformancePackName: conformancePackDetail.ConformancePackName,
				Filters: &types.ConformancePackEvaluationFilters{
					ConfigRuleNames: []string{*conformancePackRuleCompliance.ConfigRuleName},
				},
			}
			getPaginator := configservice.NewGetConformancePackComplianceDetailsPaginator(cs, detailParams)
			for getPaginator.HasMorePages() {
				getPage, err := getPaginator.NextPage(ctx, func(options *configservice.Options) {
					options.Region = cl.Region
				})
				if err != nil {
					return err
				}
				for _, conformancePackComplianceDetail := range getPage.ConformancePackRuleEvaluationResults {
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
			}
		}
	}
	return nil
}
