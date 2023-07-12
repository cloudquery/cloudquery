package efs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEfsFilesystemsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEfsClient(ctrl)
	l := types.FileSystemDescription{}
	require.NoError(t, faker.FakeObject(&l))

	m.EXPECT().DescribeFileSystems(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&efs.DescribeFileSystemsOutput{
			FileSystems: []types.FileSystemDescription{l},
		}, nil)

	b := efs.DescribeBackupPolicyOutput{}
	require.NoError(t, faker.FakeObject(&b))

	m.EXPECT().DescribeBackupPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&b, nil)

	return client.Services{
		Efs: m,
	}
}

func TestEfsFilesystems(t *testing.T) {
	client.AwsMockTestHelper(t, Filesystems(), buildEfsFilesystemsMock, client.TestOptions{})
}
