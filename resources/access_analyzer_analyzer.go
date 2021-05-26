package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func AccessAnalyzerAnalyzer() *schema.Table {
	return &schema.Table{
		Name:         "aws_access_analyzer_analyzers",
		Resolver:     fetchAccessAnalyzerAnalyzers,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "arn",
				Type: schema.TypeString,
			},
			{
				Name: "created_at",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "status",
				Type: schema.TypeString,
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
			{
				Name: "last_resource_analyzed",
				Type: schema.TypeString,
			},
			{
				Name: "last_resource_analyzed_at",
				Type: schema.TypeTimestamp,
			},
			{
				Name:     "status_reason_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatusReason.Code"),
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_access_analyzer_analyzer_findings",
				Resolver: fetchAccessAnalyzerAnalyzerFindings,
				Columns: []schema.Column{
					{
						Name:     "analyzer_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "analyzed_at",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "condition",
						Type: schema.TypeJSON,
					},
					{
						Name: "created_at",
						Type: schema.TypeTimestamp,
					},
					{
						Name:     "finding_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Id"),
					},
					{
						Name: "resource_owner_account",
						Type: schema.TypeString,
					},
					{
						Name: "resource_type",
						Type: schema.TypeString,
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
					{
						Name: "updated_at",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "action",
						Type: schema.TypeStringArray,
					},
					{
						Name: "error",
						Type: schema.TypeString,
					},
					{
						Name: "is_public",
						Type: schema.TypeBool,
					},
					{
						Name: "principal",
						Type: schema.TypeJSON,
					},
					{
						Name: "resource",
						Type: schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "aws_access_analyzer_analyzer_finding_sources",
						Resolver: fetchAccessAnalyzerAnalyzerFindingSources,
						Columns: []schema.Column{
							{
								Name:     "analyzer_finding_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "type",
								Type: schema.TypeString,
							},
							{
								Name:     "detail_access_point_arn",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Detail.AccessPointArn"),
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchAccessAnalyzerAnalyzers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	config := accessanalyzer.ListAnalyzersInput{}
	c := meta.(*client.Client)
	svc := c.Services().Analyzer
	for {
		response, err := svc.ListAnalyzers(ctx, &config, func(options *accessanalyzer.Options) {
			options.APIOptions = append(options.APIOptions, func(stack *middleware.Stack) error {
				if err := stack.Initialize.Add(&awsmiddleware.RegisterServiceMetadata{
					Region:        c.Region,
					ServiceID:     accessanalyzer.ServiceID,
					SigningName:   "access-analyzer",
					OperationName: "ListAnalyzers",
				}, middleware.Before); err != nil {
					return nil
				}
				return nil
			})
		})
		if err != nil {
			return err
		}

		res <- response.Analyzers
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func fetchAccessAnalyzerAnalyzerFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	analyzer := parent.Item.(types.AnalyzerSummary)
	c := meta.(*client.Client)
	svc := c.Services().Analyzer
	config := accessanalyzer.ListFindingsInput{
		AnalyzerArn: analyzer.Arn,
	}
	for {
		response, err := svc.ListFindings(ctx, &config, func(options *accessanalyzer.Options) {
			options.APIOptions = append(options.APIOptions, func(stack *middleware.Stack) error {
				if err := stack.Initialize.Add(&awsmiddleware.RegisterServiceMetadata{
					Region:        c.Region,
					ServiceID:     accessanalyzer.ServiceID,
					SigningName:   "access-analyzer",
					OperationName: "ListFindings",
				}, middleware.Before); err != nil {
					return nil
				}
				return nil
			})
		})
		if err != nil {
			return err
		}

		res <- response.Findings
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func fetchAccessAnalyzerAnalyzerFindingSources(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	finding := parent.Item.(types.FindingSummary)
	res <- finding.Sources
	return nil
}
