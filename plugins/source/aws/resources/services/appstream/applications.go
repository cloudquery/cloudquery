package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Applications() *schema.Table {
	tableName := "aws_appstream_applications"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Application.html`,
		Resolver:    fetchAppstreamApplications,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "appstream2"),
		Transform:   transformers.TransformWithStruct(&types.Application{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			applicationFleetAssociations(),
		},
	}
}

func fetchAppstreamApplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input appstream.DescribeApplicationsInput
	c := meta.(*client.Client)
	svc := c.Services().Appstream
	for {
		response, err := svc.DescribeApplications(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.Applications
		if response.NextToken == nil {
			break
		}
		input.NextToken = response.NextToken
	}

	return nil
}
