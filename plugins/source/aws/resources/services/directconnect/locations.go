package directconnect

import (
	"context"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Locations() *schema.Table {
	tableName := "aws_directconnect_locations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/directconnect/latest/APIReference/API_Location.html`,
		Resolver:    fetchDirectConnectLocations,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "directconnect"),
		Transform:   transformers.TransformWithStruct(&types.Location{}, transformers.WithPrimaryKeys("LocationCode")),
		Columns: []schema.Column{
			{
				Name:       "request_account_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAWSAccount,
				PrimaryKey: true,
			},
			{
				Name:       "request_region",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAWSRegion,
				PrimaryKey: true,
			},
		},
	}
}

func fetchDirectConnectLocations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config directconnect.DescribeLocationsInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDirectconnect).Directconnect
	output, err := svc.DescribeLocations(ctx, &config, func(options *directconnect.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- output.Locations
	return nil
}
