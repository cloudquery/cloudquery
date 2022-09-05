// Code generated by codegen using template resource_get.go.tpl; DO NOT EDIT.

package accessanalyzer

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
)

func AccessAnalyzerAccessanalyzersArchiveRules() *schema.Table {
	return &schema.Table{
		Name:      "aws_accessanalyzer_accessanalyzers_archive_rules",
		Resolver:  fetchAccessAnalyzerAccessanalyzersArchiveRules,
		Multiplex: client.ServiceAccountRegionMultiplexer("accessanalyzer"),
		Columns: []schema.Column{
			{
				Name:     "accessanalyzer_cq_id",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "filter",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Filter"),
			},
			{
				Name:     "rule_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RuleName"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
		},
	}
}

func fetchAccessAnalyzerAccessanalyzersArchiveRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().AccessAnalyzer

	r1 := parent.Item.(types.AnalyzerSummary)

	input := accessanalyzer.ListArchiveRulesInput{
		AnalyzerName: r1.Name,
	}

	for {
		response, err := svc.ListArchiveRules(ctx, &input)
		if err != nil {

			return diag.WrapError(err)
		}

		res <- response.ArchiveRules

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
