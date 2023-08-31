package fsx

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildFilesystemsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockFsxClient(ctrl)

	var f types.FileSystem
	require.NoError(t, faker.FakeObject(&f, faker.WithMaxDepth(5)))

	f.FileSystemType = types.FileSystemTypeLustre
	f.Lifecycle = types.FileSystemLifecycleAvailable
	f.StorageType = types.StorageTypeHdd
	m.EXPECT().DescribeFileSystems(
		gomock.Any(),
		&fsx.DescribeFileSystemsInput{MaxResults: aws.Int32(1000)},
		gomock.Any(),
	).Return(
		&fsx.DescribeFileSystemsOutput{FileSystems: []types.FileSystem{f}},
		nil,
	)

	return client.Services{
		Fsx: m,
	}
}

func TestFilesystems(t *testing.T) {
	client.AwsMockTestHelper(t, FileSystems(), buildFilesystemsMock, client.TestOptions{})
}
