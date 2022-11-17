// Code generated by codegen; DO NOT EDIT.
package resources

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

type (
	RuntimePagerArmresourcesResourceGroupsClientListResponse = runtime.Pager[armresources.ResourceGroupsClientListResponse]
)

//go:generate mockgen -package=mocks -destination=../../mocks/resources/resource_groups.go -source=resource_groups.go ResourceGroupsClient
type ResourceGroupsClient interface {
	Get(context.Context, string, *armresources.ResourceGroupsClientGetOptions) (armresources.ResourceGroupsClientGetResponse, error)
	NewListPager(*armresources.ResourceGroupsClientListOptions) *RuntimePagerArmresourcesResourceGroupsClientListResponse
}
