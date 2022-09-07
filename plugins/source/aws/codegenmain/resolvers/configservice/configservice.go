package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

func ResolveConfigRecorderArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	cfg := resource.Item.(types.ConfigurationRecorder)
	return errors.WithStack(resource.Set(c.Name, cl.ARN("config", "config-recorder", *cfg.Name)))
}

type ConformancePackComplianceWrapper struct {
	types.ConformancePackRuleCompliance
	types.ConformancePackEvaluationResult
}

func FetchConformancePackRuleCompliances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	conformancePackDetail := parent.Item.(types.ConformancePackDetail)
	cl := meta.(*client.Client)
	cs := cl.Services().ConfigService
	params := configservice.DescribeConformancePackComplianceInput{
		ConformancePackName: conformancePackDetail.ConformancePackName,
	}
	for {
		resp, err := cs.DescribeConformancePackCompliance(ctx, &params)
		if err != nil {
			return errors.WithStack(err)
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
					return errors.WithStack(err)
				}
				for _, conformancePackComplianceDetail := range output.ConformancePackRuleEvaluationResults {
					res <- ConformancePackComplianceWrapper{
						ConformancePackRuleCompliance:   conformancePackRuleCompliance,
						ConformancePackEvaluationResult: conformancePackComplianceDetail,
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
