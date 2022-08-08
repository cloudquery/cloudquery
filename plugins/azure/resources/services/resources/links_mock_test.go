package resources

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-09-01/links"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildResourceLinksMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockLinksClient(ctrl)
	var v links.ResourceLink
	if err := faker.FakeDataSkipFields(&v, []string{"Type"}); err != nil {
		t.Fatal(err)
	}
	v.Type = "sometype"
	m.EXPECT().ListAtSubscription(gomock.Any(), "").Return(
		links.NewResourceLinkResultPage(
			links.ResourceLinkResult{Value: &[]links.ResourceLink{v}},
			func(c context.Context, rlr links.ResourceLinkResult) (links.ResourceLinkResult, error) {
				return links.ResourceLinkResult{}, nil
			},
		),
		nil,
	)
	return services.Services{
		Resources: services.ResourcesClient{Links: m},
	}
}

func TestResourceLinks(t *testing.T) {
	client.AzureMockTestHelper(t, ResourcesLinks(), buildResourceLinksMock, client.TestOptions{})
}
