//go:generate mockgen -destination=./mocks/compute.go -package=mocks . DisksClient,VirtualMachinesClient,VirtualMachineExtensionsClient,VirtualMachineScaleSetsClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
	"github.com/Azure/go-autorest/autorest"
)

type ComputeClient struct {
	Disks                    DisksClient
	VirtualMachines          VirtualMachinesClient
	VirtualMachineExtensions VirtualMachineExtensionsClient
	VirtualMachineScaleSets  VirtualMachineScaleSetsClient
}

func NewComputeClient(subscriptionId string, auth autorest.Authorizer) ComputeClient {
	disks := compute.NewDisksClient(subscriptionId)
	disks.Authorizer = auth

	vmsSvc := compute.NewVirtualMachinesClient(subscriptionId)
	vmsSvc.Authorizer = auth

	vmsEx := compute.NewVirtualMachineExtensionsClient(subscriptionId)
	vmsEx.Authorizer = auth

	vmsScaleSets := compute.NewVirtualMachineScaleSetsClient(subscriptionId)
	vmsScaleSets.Authorizer = auth

	return ComputeClient{
		Disks:                    disks,
		VirtualMachines:          vmsSvc,
		VirtualMachineExtensions: vmsEx,
		VirtualMachineScaleSets:  vmsScaleSets,
	}
}

type DisksClient interface {
	List(ctx context.Context) (result compute.DiskListPage, err error)
}

type VirtualMachinesClient interface {
	ListAll(ctx context.Context, statusOnly string) (result compute.VirtualMachineListResultPage, err error)
	InstanceView(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachineInstanceView, err error)
}

type VirtualMachineExtensionsClient interface {
	List(ctx context.Context, resourceGroupName string, VMName string, expand string) (result compute.VirtualMachineExtensionsListResult, err error)
}

type VirtualMachineScaleSetsClient interface {
	ListAll(ctx context.Context) (result compute.VirtualMachineScaleSetListWithLinkResultPage, err error)
}
