// Auto generated code - DO NOT EDIT.

package compute

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
)

func TestComputeDisks(t *testing.T) {
	client.MockTestHelper(t, Disks(), createDisksMock)
}

func createDisksMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockComputeDisksClient(ctrl)
	s := services.Services{
		Compute: services.ComputeClient{
			Disks: mockClient,
		},
	}

	data := compute.Disk{}
	require.Nil(t, faker.FakeObject(&data))

	result := compute.NewDiskListPage(compute.DiskList{Value: &[]compute.Disk{data}}, func(ctx context.Context, result compute.DiskList) (compute.DiskList, error) {
		return compute.DiskList{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
