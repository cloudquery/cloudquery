package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
	"github.com/Azure/go-autorest/autorest"
)

type ComputeClient struct {
	Disks           DisksClient
	VirtualMachines VirtualMachinesClient
}

func NewComputeClient(subscriptionId string, auth autorest.Authorizer) ComputeClient {
	disks := compute.NewDisksClient(subscriptionId)
	disks.Authorizer = auth

	vmsSvc := compute.NewVirtualMachinesClient(subscriptionId)
	vmsSvc.Authorizer = auth
	return ComputeClient{
		Disks:           disks,
		VirtualMachines: vmsSvc,
	}
}

type DisksClient interface {
	List(ctx context.Context) (result compute.DiskListPage, err error)
}

type VirtualMachinesClient interface {
	ListAll(ctx context.Context, statusOnly string) (result compute.VirtualMachineListResultPage, err error)
}
