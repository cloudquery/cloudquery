//go:generate mockgen -destination=./mocks/front_door.go -package=mocks . FrontDoorClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-11-01/frontdoor"
	"github.com/Azure/go-autorest/autorest"
)

type FrontDoorClient interface {
	List(ctx context.Context) (result frontdoor.ListResultPage, err error)
}

func NewFrontDoorClient(subscriptionId string, auth autorest.Authorizer) FrontDoorClient {
	cl := frontdoor.NewFrontDoorsClient(subscriptionId)
	cl.Authorizer = auth
	return cl
}
