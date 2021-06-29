package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest"
)

type ADGroupsClient interface {
	List(ctx context.Context, filter string) (result graphrbac.GroupListResultPage, err error)
}

type ADServicePrinicpals interface {
	List(ctx context.Context, filter string) (result graphrbac.ServicePrincipalListResultPage, err error)
}

type ADUsersClient interface {
	List(ctx context.Context, filter string, expand string) (result graphrbac.UserListResultPage, err error)
}

type AD struct {
	Groups            ADGroupsClient
	ServicePrincipals ADServicePrinicpals
	Users             ADUsersClient
}

func NewADClient(subscriptionId string, auth autorest.Authorizer) AD {
	groups := graphrbac.NewGroupsClient(subscriptionId)
	groups.Authorizer = auth
	users := graphrbac.NewUsersClient(subscriptionId)
	users.Authorizer = auth
	sp := graphrbac.NewServicePrincipalsClient(subscriptionId)
	sp.Authorizer = auth
	return AD{
		Groups:            groups,
		ServicePrincipals: sp,
		Users:             users,
	}
}
