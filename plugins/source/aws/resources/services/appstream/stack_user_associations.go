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

func stackUserAssociations() *schema.Table {
	tableName := "aws_appstream_stack_user_associations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_UserStackAssociation.html`,
		Resolver:    fetchAppstreamStackUserAssociations,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "appstream2"),
		Transform:   transformers.TransformWithStruct(&types.UserStackAssociation{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "stack_name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("StackName"),
				PrimaryKey: true,
			},
			{
				Name:       "user_name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("UserName"),
				PrimaryKey: true,
			},
			{
				Name:       "authentication_type",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("AuthenticationType"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchAppstreamStackUserAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input appstream.DescribeUserStackAssociationsInput
	input.StackName = parent.Item.(types.Stack).Name
	input.MaxResults = aws.Int32(25)

	cl := meta.(*client.Client)
	svc := cl.Services().Appstream
	// No paginator available
	for {
		response, err := svc.DescribeUserStackAssociations(ctx, &input, func(options *appstream.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.UserStackAssociations
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}

	return nil
}
