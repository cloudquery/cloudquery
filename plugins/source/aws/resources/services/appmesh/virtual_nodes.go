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

func virtualNodes() *schema.Table {
	return &schema.Table{
		Name:                "aws_appmesh_virtual_nodes",
		Description:         `https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_VirtualNodeData.html`,
		Resolver:            fetchVirtualNodes,
		PreResourceResolver: getVirtualNode,
		Transform:           transformers.TransformWithStruct(&types.VirtualNodeData{}),
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

func fetchVirtualNodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAppmesh).Appmesh
	md := parent.Item.(*types.MeshData)
	input := &appmesh.ListVirtualNodesInput{
		MeshName:  md.MeshName,
		MeshOwner: md.Metadata.MeshOwner,
	}
	paginator := appmesh.NewListVirtualNodesPaginator(svc, input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *appmesh.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.VirtualNodes
	}
	return nil
}

func getVirtualNode(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAppmesh).Appmesh
	vnr := resource.Item.(types.VirtualNodeRef)
	input := appmesh.DescribeVirtualNodeInput{
		MeshName:        vnr.MeshName,
		VirtualNodeName: vnr.VirtualNodeName,
		MeshOwner:       vnr.MeshOwner,
	}
	output, err := svc.DescribeVirtualNode(ctx, &input, func(o *appmesh.Options) { o.Region = cl.Region })
	if err != nil {
		return err
	}
	resource.Item = output.VirtualNode
	return nil
}
