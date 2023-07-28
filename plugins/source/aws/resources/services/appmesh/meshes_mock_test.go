package appmesh

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildMeshes(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockAppmeshClient(ctrl)

	var mr types.MeshRef
	require.NoError(t, faker.FakeObject(&mr))

	mock.EXPECT().ListMeshes(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&appmesh.ListMeshesOutput{
			Meshes: []types.MeshRef{mr},
		},
		nil,
	)

	var md types.MeshData
	require.NoError(t, faker.FakeObject(&md))

	mock.EXPECT().DescribeMesh(
		gomock.Any(),
		&appmesh.DescribeMeshInput{MeshName: mr.MeshName, MeshOwner: mr.MeshOwner},
		gomock.Any(),
	).Return(
		&appmesh.DescribeMeshOutput{
			Mesh: &md,
		},
		nil,
	)

	// virtual gateways
	var vgr types.VirtualGatewayRef
	require.NoError(t, faker.FakeObject(&vgr))

	mock.EXPECT().ListVirtualGateways(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&appmesh.ListVirtualGatewaysOutput{
			VirtualGateways: []types.VirtualGatewayRef{vgr},
		},
		nil,
	)

	var vgd types.VirtualGatewayData
	require.NoError(t, faker.FakeObject(&vgd))

	mock.EXPECT().DescribeVirtualGateway(
		gomock.Any(),
		&appmesh.DescribeVirtualGatewayInput{
			MeshName:           mr.MeshName,
			MeshOwner:          mr.MeshOwner,
			VirtualGatewayName: vgr.VirtualGatewayName,
		},
		gomock.Any(),
	).Return(
		&appmesh.DescribeVirtualGatewayOutput{
			VirtualGateway: &vgd,
		},
		nil,
	)

	// virtual nodes
	var vnr types.VirtualNodeRef
	require.NoError(t, faker.FakeObject(&vnr))

	mock.EXPECT().ListVirtualNodes(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&appmesh.ListVirtualNodesOutput{
			VirtualNodes: []types.VirtualNodeRef{vnr},
		},
		nil,
	)

	var vnd types.VirtualNodeData
	require.NoError(t, faker.FakeObject(&vnd))

	mock.EXPECT().DescribeVirtualNode(
		gomock.Any(),
		&appmesh.DescribeVirtualNodeInput{
			MeshName:        mr.MeshName,
			MeshOwner:       mr.MeshOwner,
			VirtualNodeName: vnr.VirtualNodeName},
		gomock.Any(),
	).Return(
		&appmesh.DescribeVirtualNodeOutput{
			VirtualNode: &vnd,
		},
		nil,
	)

	// virtual services
	var vsr types.VirtualServiceRef
	require.NoError(t, faker.FakeObject(&vsr))

	mock.EXPECT().ListVirtualServices(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&appmesh.ListVirtualServicesOutput{
			VirtualServices: []types.VirtualServiceRef{vsr},
		},
		nil,
	)

	var vsd types.VirtualServiceData
	require.NoError(t, faker.FakeObject(&vsd))

	mock.EXPECT().DescribeVirtualService(
		gomock.Any(),
		&appmesh.DescribeVirtualServiceInput{
			MeshName:           mr.MeshName,
			MeshOwner:          mr.MeshOwner,
			VirtualServiceName: vsr.VirtualServiceName,
		},
		gomock.Any(),
	).Return(
		&appmesh.DescribeVirtualServiceOutput{
			VirtualService: &vsd,
		},
		nil,
	)

	// virtual routers
	var vrr types.VirtualRouterRef
	require.NoError(t, faker.FakeObject(&vrr))

	mock.EXPECT().ListVirtualRouters(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&appmesh.ListVirtualRoutersOutput{
			VirtualRouters: []types.VirtualRouterRef{vrr},
		},
		nil,
	)

	var vrd types.VirtualRouterData
	require.NoError(t, faker.FakeObject(&vrd))

	mock.EXPECT().DescribeVirtualRouter(
		gomock.Any(),
		&appmesh.DescribeVirtualRouterInput{
			MeshName:          mr.MeshName,
			MeshOwner:         mr.MeshOwner,
			VirtualRouterName: vrr.VirtualRouterName,
		},
		gomock.Any(),
	).Return(
		&appmesh.DescribeVirtualRouterOutput{
			VirtualRouter: &vrd,
		},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&appmesh.ListTagsForResourceInput{ResourceArn: md.Metadata.Arn},
		gomock.Any(),
	).Return(
		&appmesh.ListTagsForResourceOutput{
			Tags: []types.TagRef{
				{Key: aws.String("key"), Value: aws.String("value")},
			},
		},
		nil,
	)
	return client.Services{Appmesh: mock}
}

func TestMeshes(t *testing.T) {
	client.AwsMockTestHelper(t, Meshes(), buildMeshes, client.TestOptions{})
}
