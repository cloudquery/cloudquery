// Code generated by codegen; DO NOT EDIT.
package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

type (
	RuntimePagerArmcomputeDisksClientListByResourceGroupResponse = runtime.Pager[armcompute.DisksClientListByResourceGroupResponse]
	RuntimePagerArmcomputeDisksClientListResponse                = runtime.Pager[armcompute.DisksClientListResponse]
)

//go:generate mockgen -package=mocks -destination=../../mocks/compute/disks.go -source=disks.go DisksClient
type DisksClient interface {
	Get(context.Context, string, string, *armcompute.DisksClientGetOptions) (armcompute.DisksClientGetResponse, error)
	NewListByResourceGroupPager(string, *armcompute.DisksClientListByResourceGroupOptions) *RuntimePagerArmcomputeDisksClientListByResourceGroupResponse
	NewListPager(*armcompute.DisksClientListOptions) *RuntimePagerArmcomputeDisksClientListResponse
}
