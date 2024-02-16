package appmesh

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func virtualGateways() *schema.Table {
	return &schema.Table{
		Name:                "aws_appmesh_virtual_gateways",
		Description:         `https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_VirtualGatewayData.html`,
		Resolver:            fetchVirtualGateways,
		PreResourceResolver: getVirtualGateway,
		Transform:           transformers.TransformWithStruct(&types.VirtualGatewayData{}),
		Columns: []schema.Column{
			client.RequestAccountIDColumn(true),
			client.RequestRegionColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Metadata.Arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "mesh_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchVirtualGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAppmesh).Appmesh
	md := parent.Item.(*types.MeshData)
	input := &appmesh.ListVirtualGatewaysInput{
		MeshName:  md.MeshName,
		MeshOwner: md.Metadata.MeshOwner,
	}
	paginator := appmesh.NewListVirtualGatewaysPaginator(svc, input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *appmesh.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.VirtualGateways
	}
	return nil
}

func getVirtualGateway(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAppmesh).Appmesh
	vgr := resource.Item.(types.VirtualGatewayRef)
	input := appmesh.DescribeVirtualGatewayInput{
		MeshName:           vgr.MeshName,
		VirtualGatewayName: vgr.VirtualGatewayName,
		MeshOwner:          vgr.MeshOwner,
	}
	output, err := svc.DescribeVirtualGateway(ctx, &input, func(o *appmesh.Options) { o.Region = cl.Region })
	if err != nil {
		return err
	}
	resource.Item = output.VirtualGateway
	return nil
}
