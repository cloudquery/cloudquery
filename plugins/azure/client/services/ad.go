package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest"
	auth2 "github.com/Azure/go-autorest/autorest/azure/auth"
)

type ADApplicationsClient interface {
	List(ctx context.Context, filter string) (result graphrbac.ApplicationListResultPage, err error)
}

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
	Applications      ADApplicationsClient
	Groups            ADGroupsClient
	ServicePrincipals ADServicePrinicpals
	Users             ADUsersClient
}

func NewADClient(_ string, _ autorest.Authorizer) AD {
	//ad services need custom authorization
	auth, _ := auth2.NewAuthorizerFromEnvironmentWithResource(graphrbac.DefaultBaseURI) //https://graph.windows.net

	//we ignore errors since operations below usually are already executed during creation of a default authorizer
	settings, _ := auth2.GetSettingsFromEnvironment()
	c, _ := settings.GetClientCredentials()

	apps := graphrbac.NewApplicationsClient(c.TenantID)
	apps.Authorizer = auth
	groups := graphrbac.NewGroupsClient(c.TenantID)
	groups.Authorizer = auth
	users := graphrbac.NewUsersClient(c.TenantID)
	users.Authorizer = auth
	sp := graphrbac.NewServicePrincipalsClient(c.TenantID)
	sp.Authorizer = auth
	return AD{
		Applications:      apps,
		Groups:            groups,
		ServicePrincipals: sp,
		Users:             users,
	}
}
