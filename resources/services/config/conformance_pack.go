package config

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

type ConformancePackComplianceWrapper struct {
	types.ConformancePackRuleCompliance
	types.ConformancePackEvaluationResult
}

func ConfigConformancePack() *schema.Table {
	return &schema.Table{
		Name:          "aws_config_conformance_packs",
		Description:   "Returns details of a conformance pack.",
		Resolver:      fetchConfigConformancePacks,
		Multiplex:     client.ServiceAccountRegionMultiplexer("config"),
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "Amazon Resource Name (ARN) of the conformance pack.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ConformancePackArn"),
			},
			{
				Name:        "conformance_pack_id",
				Description: "ID of the conformance pack.",
				Type:        schema.TypeString,
			},
			{
				Name:        "conformance_pack_name",
				Description: "Name of the conformance pack.",
				Type:        schema.TypeString,
			},
			{
				Name:        "conformance_pack_input_parameters",
				Description: "A list of ConformancePackInputParameter objects.",
				Type:        schema.TypeJSON,
				Resolver:    resolveConfigConformancePackConformancePackInputParameters,
			},
			{
				Name:        "created_by",
				Description: "AWS service that created the conformance pack.",
				Type:        schema.TypeString,
			},
			{
				Name:        "delivery_s3_bucket",
				Description: "Amazon S3 bucket where AWS Config stores conformance pack templates.",
				Type:        schema.TypeString,
			},
			{
				Name:        "delivery_s3_key_prefix",
				Description: "The prefix for the Amazon S3 bucket.",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_update_requested_time",
				Description: "Last time when conformation pack update was requested.",
				Type:        schema.TypeTimestamp,
			},
		},
		Relations: []*schema.Table{
			{
				Name:         "aws_config_conformance_pack_rule_compliances",
				Description:  "Compliance information of one or more AWS Config rules within a conformance pack",
				Resolver:     fetchConfigConformancePackRuleCompliances,
				Multiplex:    client.ServiceAccountRegionMultiplexer("config"),
				IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
				DeleteFilter: client.DeleteAccountRegionFilter,
				Columns: []schema.Column{
					{
						Name:        "conformance_pack_cq_id",
						Description: "Unique CloudQuery ID of aws_config_conformance_packs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "compliance_type",
						Description: "Compliance of the AWS Config rule",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ConformancePackRuleCompliance.ComplianceType"),
					},
					{
						Name:        "config_rule_name",
						Description: "Name of the config rule.",
						Type:        schema.TypeString,
					},
					{
						Name:        "controls",
						Description: "Controls for the conformance pack",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "config_rule_invoked_time",
						Description: "The time when AWS Config rule evaluated AWS resource.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "resource_id",
						Description: "The ID of the evaluated AWS resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EvaluationResultIdentifier.EvaluationResultQualifier.ResourceId"),
					},
					{
						Name:        "resource_type",
						Description: "The type of AWS resource that was evaluated.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EvaluationResultIdentifier.EvaluationResultQualifier.ResourceType"),
					},
					{
						Name:        "ordering_timestamp",
						Description: "The time of the event that triggered the evaluation of your AWS resources.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("EvaluationResultIdentifier.OrderingTimestamp"),
					},
					{
						Name:        "result_recorded_time",
						Description: "The time when AWS Config recorded the evaluation result.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "annotation",
						Description: "Supplementary information about how the evaluation determined the compliance.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchConfigConformancePacks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	config := configservice.DescribeConformancePacksInput{}
	var ae smithy.APIError
	for {
		resp, err := c.Services().ConfigService.DescribeConformancePacks(ctx, &config, func(options *configservice.Options) {
			options.Region = c.Region
		})

		// This is a workaround until this bug is fixed = https://github.com/aws/aws-sdk-go-v2/issues/1539
		if (c.Region == "af-south-1" || c.Region == "ap-northeast-3") && errors.As(err, &ae) && ae.ErrorCode() == "AccessDeniedException" {
			return nil
		}

		if err != nil {
			return diag.WrapError(err)
		}
		res <- resp.ConformancePackDetails
		if resp.NextToken == nil {
			break
		}
		config.NextToken = resp.NextToken
	}
	return nil
}
func resolveConfigConformancePackConformancePackInputParameters(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	conformancePack := resource.Item.(types.ConformancePackDetail)
	params := make(map[string]*string, len(conformancePack.ConformancePackInputParameters))
	for _, p := range conformancePack.ConformancePackInputParameters {
		params[*p.ParameterName] = p.ParameterValue
	}
	return diag.WrapError(resource.Set(c.Name, params))
}

func fetchConfigConformancePackRuleCompliances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	conformancePackDetail := parent.Item.(types.ConformancePackDetail)
	c := meta.(*client.Client)
	cs := c.Services().ConfigService
	params := configservice.DescribeConformancePackComplianceInput{
		ConformancePackName: conformancePackDetail.ConformancePackName,
	}
	for {
		resp, err := cs.DescribeConformancePackCompliance(ctx, &params, func(options *configservice.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, conformancePackRuleCompliance := range resp.ConformancePackRuleComplianceList {
			detailParams := &configservice.GetConformancePackComplianceDetailsInput{
				ConformancePackName: conformancePackDetail.ConformancePackName,
				Filters: &types.ConformancePackEvaluationFilters{
					ConfigRuleNames: []string{*conformancePackRuleCompliance.ConfigRuleName},
				},
			}
			for {
				output, err := cs.GetConformancePackComplianceDetails(ctx, detailParams, func(options *configservice.Options) {
					options.Region = c.Region
				})
				if err != nil {
					return diag.WrapError(err)
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
