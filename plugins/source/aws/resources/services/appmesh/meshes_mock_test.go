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
