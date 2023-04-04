package inspector2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Findings() *schema.Table {
	tableName := "aws_inspector2_findings"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/inspector/v2/APIReference/API_Finding.html`,
		Resolver:    fetchInspector2Findings,
		Transform:   transformers.TransformWithStruct(&types.Finding{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "inspector2"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FindingArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchInspector2Findings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Inspector2
	input := inspector2.ListFindingsInput{MaxResults: aws.Int32(100)}
	for {
		response, err := svc.ListFindings(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.Findings
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
