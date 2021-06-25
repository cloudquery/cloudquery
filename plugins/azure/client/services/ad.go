package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest"
)

type ADUsersClient interface {
	List(ctx context.Context, filter string, expand string) (result graphrbac.UserListResultPage, err error)
}

type AD struct {
	Users ADUsersClient
}

func NewADClient(subscriptionId string, auth autorest.Authorizer) AD {
	client := graphrbac.NewUsersClient(subscriptionId)
	client.Authorizer = auth
	return AD{Users: client}
}
