// Code generated by codegen; DO NOT EDIT.
package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

type (
	RuntimePagerArmsqlVirtualNetworkRulesClientListByServerResponse = runtime.Pager[armsql.VirtualNetworkRulesClientListByServerResponse]
)

//go:generate mockgen -package=mocks -destination=../../mocks/sql/virtual_network_rules.go -source=virtual_network_rules.go VirtualNetworkRulesClient
type VirtualNetworkRulesClient interface {
	Get(context.Context, string, string, string, *armsql.VirtualNetworkRulesClientGetOptions) (armsql.VirtualNetworkRulesClientGetResponse, error)
	NewListByServerPager(string, string, *armsql.VirtualNetworkRulesClientListByServerOptions) *RuntimePagerArmsqlVirtualNetworkRulesClientListByServerResponse
}
