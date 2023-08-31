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

func buildStorageVmsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockFsxClient(ctrl)

	var vm types.StorageVirtualMachine
	require.NoError(t, faker.FakeObject(&vm))
	m.EXPECT().DescribeStorageVirtualMachines(
		gomock.Any(),
		&fsx.DescribeStorageVirtualMachinesInput{MaxResults: aws.Int32(1000)},
		gomock.Any(),
	).Return(
		&fsx.DescribeStorageVirtualMachinesOutput{StorageVirtualMachines: []types.StorageVirtualMachine{vm}},
		nil,
	)
	return client.Services{
		Fsx: m,
	}
}

func TestStorageVms(t *testing.T) {
	client.AwsMockTestHelper(t, StorageVirtualMachines(), buildStorageVmsMock, client.TestOptions{})
}
