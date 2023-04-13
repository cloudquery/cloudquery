package inspector2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Findings() *schema.Table {
	tableName := "aws_inspector2_findings"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/inspector/v2/APIReference/API_Finding.html
The 'request_account_id' and 'request_region' columns are added to show from where the request was made.`,
		Resolver:  fetchInspector2Findings,
		Transform: transformers.TransformWithStruct(&types.Finding{}),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "inspector2"),
		Columns: []schema.Column{
			{
				Name:            "request_account_id",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "request_region",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
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
	paginator := inspector2.NewListFindingsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Findings
	}
	return nil
}
