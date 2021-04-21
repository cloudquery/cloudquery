package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/2020-09-01/network/mgmt/network"
	"github.com/Azure/go-autorest/autorest"
)

type NetworksClient struct {
	VirtualNetworks VirtualNetworksClient
}

func NewNetworksClient(subscriptionId string, auth autorest.Authorizer) NetworksClient {
	vn := network.NewVirtualNetworksClient(subscriptionId)
	vn.Authorizer = auth
	return NetworksClient{
		VirtualNetworks: vn,
	}
}

type VirtualNetworksClient interface {
	ListAll(ctx context.Context) (result network.VirtualNetworkListResultPage, err error)
}
