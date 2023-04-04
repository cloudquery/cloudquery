package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Findings() *schema.Table {
	tableName := "aws_inspector_findings"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/inspector/v1/APIReference/API_Finding.html`,
		Resolver:    fetchInspectorFindings,
		Transform:   transformers.TransformWithStruct(&types.Finding{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "inspector"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "attributes",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTagField("Attributes"),
			},
			{
				Name:     "user_attributes",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTagField("UserAttributes"),
			},
		},
	}
}

func fetchInspectorFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Inspector
	input := inspector.ListFindingsInput{MaxResults: aws.Int32(500)}
	for {
		response, err := svc.ListFindings(ctx, &input)
		if err != nil {
			return err
		}
		if len(response.FindingArns) > 0 {
			out, err := svc.DescribeFindings(ctx, &inspector.DescribeFindingsInput{FindingArns: response.FindingArns})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return err
			}
			res <- out.Findings
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
