//go:generate mockgen -destination=./mocks/iothub.go -package=mocks . IotHubDevicesClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/iothub/mgmt/2021-07-02/devices"
	"github.com/Azure/go-autorest/autorest"
)

type IotHubClient struct {
	Devices IotHubDevicesClient
}

type IotHubDevicesClient interface {
	ListBySubscription(ctx context.Context) (result devices.IotHubDescriptionListResultPage, err error)
}

func NewIotHubClient(subscriptionId string, auth autorest.Authorizer) IotHubClient {
	cl := devices.NewIotHubResourceClient(subscriptionId)
	cl.Authorizer = auth
	return IotHubClient{
		Devices: cl,
	}
}
