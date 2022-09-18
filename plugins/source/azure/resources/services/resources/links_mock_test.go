// Auto generated code - DO NOT EDIT.

package resources

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-09-01/links"
)

func TestResourcesLinks(t *testing.T) {
	client.MockTestHelper(t, Links(), createLinksMock)
}

func createLinksMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockResourcesLinksClient(ctrl)
	s := services.Services{
		Resources: services.ResourcesClient{
			Links: mockClient,
		},
	}

	data := links.ResourceLink{}
	require.Nil(t, faker.FakeObject(&data))

	result := links.NewResourceLinkResultPage(links.ResourceLinkResult{Value: &[]links.ResourceLink{data}}, func(ctx context.Context, result links.ResourceLinkResult) (links.ResourceLinkResult, error) {
		return links.ResourceLinkResult{}, nil
	})

	mockClient.EXPECT().ListAtSubscription(gomock.Any(), "").Return(result, nil)
	return s
}
