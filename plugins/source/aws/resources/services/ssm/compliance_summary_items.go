package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ComplianceSummaryItems() *schema.Table {
	tableName := "aws_ssm_compliance_summary_items"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ComplianceSummaryItem.html`,
		Resolver:    fetchSsmComplianceSummaryItems,
		Transform:   transformers.TransformWithStruct(&types.ComplianceSummaryItem{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ssm"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name: "compliance_type",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchSsmComplianceSummaryItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Ssm

	params := ssm.ListComplianceSummariesInput{
		MaxResults: aws.Int32(50),
	}
	for {
		output, err := svc.ListComplianceSummaries(ctx, &params)
		if err != nil {
			return err
		}
		res <- output.ComplianceSummaryItems

		if aws.ToString(output.NextToken) == "" {
			break
		}
		params.NextToken = output.NextToken
	}
	return nil
}
