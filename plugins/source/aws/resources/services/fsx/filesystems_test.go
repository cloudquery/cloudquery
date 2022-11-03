package fsx

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildFilesystemsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockFsxClient(ctrl)

	var f types.FileSystem
	err := faker.FakeObject(&f, faker.WithMaxDepth(5))
	if err != nil {
		t.Fatalf("FakeObject returned error: %v", err)
	}
	f.FileSystemType = types.FileSystemTypeLustre
	f.Lifecycle = types.FileSystemLifecycleAvailable
	f.StorageType = types.StorageTypeHdd
	m.EXPECT().DescribeFileSystems(
		gomock.Any(),
		&fsx.DescribeFileSystemsInput{MaxResults: aws.Int32(1000)},
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
