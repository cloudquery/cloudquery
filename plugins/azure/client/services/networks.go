package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-08-01/network"
	"github.com/Azure/go-autorest/autorest"
)

type NetworksClient struct {
	VirtualNetworks VirtualNetworksClient
	SecurityGroups  SecurityGroupsClient
	Watchers        WatchersClient
}

func NewNetworksClient(subscriptionId string, auth autorest.Authorizer) NetworksClient {
	vn := network.NewVirtualNetworksClient(subscriptionId)
	vn.Authorizer = auth

	sg := network.NewSecurityGroupsClient(subscriptionId)
	sg.Authorizer = auth
	wch := network.NewWatchersClient(subscriptionId)
	wch.Authorizer = auth
	return NetworksClient{
		VirtualNetworks: vn,
		SecurityGroups:  sg,
		Watchers:        wch,
	}
}

type VirtualNetworksClient interface {
	ListAll(ctx context.Context) (result network.VirtualNetworkListResultPage, err error)
}

type SecurityGroupsClient interface {
	ListAll(ctx context.Context) (result network.SecurityGroupListResultPage, err error)
}

type WatchersClient interface {
	ListAll(ctx context.Context) (result network.WatcherListResult, err error)
	GetFlowLogStatus(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.FlowLogStatusParameters) (result network.WatchersGetFlowLogStatusFuture, err error)
}
