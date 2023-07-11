package appstream

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func applicationFleetAssociations() *schema.Table {
	tableName := "aws_appstream_application_fleet_associations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_ApplicationFleetAssociation.html`,
		Resolver:    fetchAppstreamApplicationFleetAssociations,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "appstream2"),
		Transform:   transformers.TransformWithStruct(&types.ApplicationFleetAssociation{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "application_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ApplicationArn"),
				PrimaryKey: true,
			},
			{
				Name:       "fleet_name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("FleetName"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchAppstreamApplicationFleetAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	parentApplication := parent.Item.(types.Application)

	var input appstream.DescribeApplicationFleetAssociationsInput
	input.ApplicationArn = parentApplication.Arn

	cl := meta.(*client.Client)
	svc := cl.Services().Appstream
	// No paginator available
	for {
		response, err := svc.DescribeApplicationFleetAssociations(ctx, &input, func(options *appstream.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.ApplicationFleetAssociations
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}

	return nil
}
