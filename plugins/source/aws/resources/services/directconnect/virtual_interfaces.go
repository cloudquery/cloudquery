package directconnect

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func VirtualInterfaces() *schema.Table {
	tableName := "aws_directconnect_virtual_interfaces"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/directconnect/latest/APIReference/API_VirtualInterface.html`,
		Resolver:    fetchDirectconnectVirtualInterfaces,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "directconnect"),
		Transform:   transformers.TransformWithStruct(&types.VirtualInterface{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveVirtualInterfaceARN(),
				PrimaryKey: true,
			},
			{
				Name:     "id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("VirtualInterfaceId"),
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchDirectconnectVirtualInterfaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config directconnect.DescribeVirtualInterfacesInput
	cl := meta.(*client.Client)
	svc := cl.Services().Directconnect
	output, err := svc.DescribeVirtualInterfaces(ctx, &config, func(options *directconnect.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- output.VirtualInterfaces
	return nil
}

func resolveVirtualInterfaceARN() schema.ColumnResolver {
	return client.ResolveARN(client.DirectConnectService, func(resource *schema.Resource) ([]string, error) {
		return []string{"dxvif", *resource.Item.(types.VirtualInterface).VirtualInterfaceId}, nil
	})
}
