// Code generated by codegen; DO NOT EDIT.

package cosmosdb

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	api "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/azure/client/mocks/cosmos"
	service "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/cosmos"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildMongoDbDatabases(t *testing.T, ctrl *gomock.Controller, c *client.Services) {
	if c.Cosmos == nil {
		c.Cosmos = new(service.CosmosClient)
	}
	cosmosClient := c.Cosmos
	if cosmosClient.MongoDBResourcesClient == nil {
		cosmosClient.MongoDBResourcesClient = mocks.NewMockMongoDBResourcesClient(ctrl)
	}

	mockMongoDBResourcesClient := cosmosClient.MongoDBResourcesClient.(*mocks.MockMongoDBResourcesClient)

	var response api.MongoDBResourcesClientListMongoDBDatabasesResponse
	require.NoError(t, faker.FakeObject(&response))
	// Use correct Azure ID format
	const id = "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	response.Value[0].ID = to.Ptr(id)

	mockMongoDBResourcesClient.EXPECT().NewListMongoDBDatabasesPager(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(client.CreatePager(response)).MinTimes(1)
}
