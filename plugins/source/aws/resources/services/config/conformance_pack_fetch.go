package config

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type ConformancePackComplianceWrapper struct {
	// Fields from types.ConformancePackRuleCompliance:

	// Compliance of the Config rule. The allowed values are COMPLIANT, NON_COMPLIANT,
	// and INSUFFICIENT_DATA.
	ComplianceType types.ConformancePackComplianceType

	// Name of the Config rule.
	ConfigRuleName *string

	// Controls for the conformance pack. A control is a process to prevent or detect
	// problems while meeting objectives. A control can align with a specific
	// compliance regime or map to internal controls defined by an organization.
	Controls []string

	// Fields from types.ConformancePackEvaluationResult:

	// The time when Config rule evaluated Amazon Web Services resource.
	//
	// This member is required.
	ConfigRuleInvokedTime *time.Time

	// Uniquely identifies an evaluation result.
	//
	// This member is required.
	EvaluationResultIdentifier *types.EvaluationResultIdentifier

	// The time when Config recorded the evaluation result.
	//
	// This member is required.
	ResultRecordedTime *time.Time

	// Supplementary information about how the evaluation determined the compliance.
	Annotation *string
}

func fetchConfigConformancePacks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	config := configservice.DescribeConformancePacksInput{}
	var ae smithy.APIError
	for {
		resp, err := c.Services().ConfigService.DescribeConformancePacks(ctx, &config)

		// This is a workaround until this bug is fixed = https://github.com/aws/aws-sdk-go-v2/issues/1539
		if (c.Region == "af-south-1" || c.Region == "ap-northeast-3") && errors.As(err, &ae) && ae.ErrorCode() == "AccessDeniedException" {
			return nil
		}

		if err != nil {
			return err
		}
		res <- resp.ConformancePackDetails
		if resp.NextToken == nil {
			break
		}
		config.NextToken = resp.NextToken
	}
	return nil
}

func fetchConfigConformancePackRuleCompliances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	conformancePackDetail := parent.Item.(types.ConformancePackDetail)
	c := meta.(*client.Client)
	cs := c.Services().ConfigService
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
					res <- ConformancePackComplianceWrapper{
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
