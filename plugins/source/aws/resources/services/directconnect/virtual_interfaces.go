package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveVirtualInterfaceARN(),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualInterfaceId"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchDirectconnectVirtualInterfaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config directconnect.DescribeVirtualInterfacesInput
	c := meta.(*client.Client)
	svc := c.Services().Directconnect
	output, err := svc.DescribeVirtualInterfaces(ctx, &config)
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
