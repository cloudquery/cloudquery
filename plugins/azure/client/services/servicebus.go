//go:generate mockgen -destination=./mocks/servicebus.go -package=mocks . NamespacesClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"
	"github.com/Azure/go-autorest/autorest"
)

type ServicebusClient struct {
	Namespaces NamespacesClient
}

type NamespacesClient interface {
	List(ctx context.Context) (result servicebus.SBNamespaceListResultPage, err error)
}

func NewServicebusClient(subscriptionID string, auth autorest.Authorizer) ServicebusClient {
	n := servicebus.NewNamespacesClient(subscriptionID)
	n.Authorizer = auth
	return ServicebusClient{
		Namespaces: n,
	}
}
