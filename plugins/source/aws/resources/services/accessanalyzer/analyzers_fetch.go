package accessanalyzer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchAccessanalyzerAnalyzers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := accessanalyzer.ListAnalyzersInput{}
	c := meta.(*client.Client)
	svc := c.Services().Accessanalyzer
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

func fetchAccessanalyzerAnalyzerFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	analyzer := parent.Item.(types.AnalyzerSummary)
	c := meta.(*client.Client)
	svc := c.Services().Accessanalyzer
	config := accessanalyzer.ListFindingsInput{
		AnalyzerArn: analyzer.Arn,
	}
	for {
		response, err := svc.ListFindings(ctx, &config)
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

func fetchAccessanalyzerAnalyzerArchiveRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	analyzer := parent.Item.(types.AnalyzerSummary)
	c := meta.(*client.Client)
	svc := c.Services().Accessanalyzer
	config := accessanalyzer.ListArchiveRulesInput{
		AnalyzerName: analyzer.Name,
	}
	for {
		response, err := svc.ListArchiveRules(ctx, &config)
		if err != nil {
			return err
		}

		res <- response.ArchiveRules
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func resolveFindingArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "accessanalyzer",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "finding_summary/" + aws.ToString(resource.Item.(types.FindingSummary).Id),
	}
	return resource.Set(c.Name, a.String())
}
