//go:build mock
// +build mock

package efs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEfsFilesystemsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEfsClient(ctrl)
	l := types.FileSystemDescription{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeFileSystems(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&efs.DescribeFileSystemsOutput{
			FileSystems: []types.FileSystemDescription{l},
		}, nil)
	return client.Services{
		EFS: m,
	}
}

func TestEfsFilesystems(t *testing.T) {
	client.AwsMockTestHelper(t, EfsFilesystems(), buildEfsFilesystemsMock, client.TestOptions{})
}
