package resources_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-01/compute"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildComputeDiskMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockDisksClient(ctrl)
	s := services.Services{
		Compute: services.ComputeClient{Disks: m},
	}
	l := compute.Disk{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	diskListPage := compute.NewDiskListPage(compute.DiskList{Value: &[]compute.Disk{l}}, func(ctx context.Context, list compute.DiskList) (compute.DiskList, error) {
		return compute.DiskList{}, nil
	},
	)
	m.EXPECT().List(gomock.Any()).Return(diskListPage, nil)
	return s
}

func TestComputeDisks(t *testing.T) {
	azureTestHelper(t, resources.ComputeDisks(), buildComputeDiskMock)
}
