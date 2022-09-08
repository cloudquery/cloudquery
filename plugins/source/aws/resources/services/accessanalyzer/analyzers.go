package accessanalyzer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/smithy-go/middleware"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Analyzers() *schema.Table {
	return &schema.Table{
		Name:        "aws_access_analyzer_analyzers",
		Description: "Contains information about the analyzer",
		Resolver:    fetchAccessAnalyzerAnalyzers,
		Multiplex:   client.ServiceAccountRegionMultiplexer("access-analyzer"),
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
				Name:            "arn",
				Description:     "The ARN of the analyzer",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "created_at",
				Description: "A timestamp for the time at which the analyzer was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "name",
				Description: "The name of the analyzer",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of the analyzer",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of analyzer, which corresponds to the zone of trust chosen for the analyzer",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_resource_analyzed",
				Description: "The resource that was most recently analyzed by the analyzer",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_resource_analyzed_at",
				Description: "The time at which the most recently analyzed resource was analyzed",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "status_reason_code",
				Description: "The reason code for the current status of the analyzer",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StatusReason.Code"),
			},
			{
				Name:        "tags",
				Description: "The tags added to the analyzer",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "findings",
				Description: "Contains information about a finding",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "archive_rules",
				Description: "Contains information about an archive rule",
				Type:        schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchAccessAnalyzerAnalyzers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
