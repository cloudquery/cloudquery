// Code generated by codegen; DO NOT EDIT.
package postgresql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

type (
	RuntimePagerArmpostgresqlServersClientListByResourceGroupResponse = runtime.Pager[armpostgresql.ServersClientListByResourceGroupResponse]
	RuntimePagerArmpostgresqlServersClientListResponse                = runtime.Pager[armpostgresql.ServersClientListResponse]
)

//go:generate mockgen -package=mocks -destination=../../mocks/postgresql/servers.go -source=servers.go ServersClient
type ServersClient interface {
	Get(context.Context, string, string, *armpostgresql.ServersClientGetOptions) (armpostgresql.ServersClientGetResponse, error)
	NewListByResourceGroupPager(string, *armpostgresql.ServersClientListByResourceGroupOptions) *RuntimePagerArmpostgresqlServersClientListByResourceGroupResponse
	NewListPager(*armpostgresql.ServersClientListOptions) *RuntimePagerArmpostgresqlServersClientListResponse
}
