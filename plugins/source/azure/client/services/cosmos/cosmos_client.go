// Code generated by codegen; DO NOT EDIT.
package cosmos

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

type CosmosClient struct {
	DatabaseAccountsClient DatabaseAccountsClient
	MongoDBResourcesClient MongoDBResourcesClient
	SQLResourcesClient     SQLResourcesClient
}

func NewCosmosClient(subscriptionID string, credentials azcore.TokenCredential, options *arm.ClientOptions) (*CosmosClient, error) {
	var client CosmosClient
	var err error

	client.DatabaseAccountsClient, err = armcosmos.NewDatabaseAccountsClient(subscriptionID, credentials, options)
	if err != nil {
		return nil, err
	}

	client.MongoDBResourcesClient, err = armcosmos.NewMongoDBResourcesClient(subscriptionID, credentials, options)
	if err != nil {
		return nil, err
	}

	client.SQLResourcesClient, err = armcosmos.NewSQLResourcesClient(subscriptionID, credentials, options)
	if err != nil {
		return nil, err
	}

	return &client, nil
}
