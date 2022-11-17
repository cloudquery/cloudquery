// Code generated by codegen; DO NOT EDIT.
package subscription

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
)

type (
	RuntimePagerArmsubscriptionTenantsClientListResponse = runtime.Pager[armsubscription.TenantsClientListResponse]
)

//go:generate mockgen -package=mocks -destination=../../mocks/subscription/tenants.go -source=tenants.go TenantsClient
type TenantsClient interface {
	NewListPager(*armsubscription.TenantsClientListOptions) *RuntimePagerArmsubscriptionTenantsClientListResponse
}
