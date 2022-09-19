//go:generate mockgen -destination=./mocks/servicebus.go -package=mocks . ServicebusNamespacesClient,ServicebusTopicsClient,ServicebusAuthorizationRulesClient,ServicebusAccessKeysClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"
	"github.com/Azure/go-autorest/autorest"
)

type ServicebusClient struct {
	Namespaces         ServicebusNamespacesClient
	Topics             ServicebusTopicsClient
	AuthorizationRules ServicebusAuthorizationRulesClient
	AccessKeys         ServicebusAccessKeysClient
}

type ServicebusNamespacesClient interface {
	List(ctx context.Context) (result servicebus.SBNamespaceListResultPage, err error)
}

type ServicebusTopicsClient interface {
	ListByNamespace(ctx context.Context, resourceGroupName string, namespaceName string, skip *int32, top *int32) (result servicebus.SBTopicListResultPage, err error)
}

type ServicebusAuthorizationRulesClient interface {
	ListAuthorizationRules(ctx context.Context, resourceGroupName string, namespaceName string, topicName string) (result servicebus.SBAuthorizationRuleListResultPage, err error)
}

type ServicebusAccessKeysClient interface {
	ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string) (result servicebus.AccessKeys, err error)
}

func NewServicebusClient(subscriptionID string, auth autorest.Authorizer) ServicebusClient {
	n := servicebus.NewNamespacesClient(subscriptionID)
	n.Authorizer = auth

	t := servicebus.NewTopicsClient(subscriptionID)
	t.Authorizer = auth
	return ServicebusClient{
		Namespaces:         n,
		Topics:             t,
		AuthorizationRules: t,
		AccessKeys:         t,
	}
}
