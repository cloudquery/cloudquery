package search

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/search/mgmt/2020-08-01/search"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSearchServices(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockSearchServiceClient(ctrl)
	var searchService search.Service
	if err := faker.FakeData(&searchService); err != nil {
		t.Fatal(err)
	}
	ip := "8.8.8.8"
	searchService.NetworkRuleSet = &search.NetworkRuleSet{
		IPRules: &[]search.IPRule{
			{
				Value: &ip,
			},
		},
	}
	m.EXPECT().ListBySubscription(gomock.Any(), nil).Return(
		search.NewServiceListResultPage(
			search.ServiceListResult{Value: &[]search.Service{searchService}},
			func(c context.Context, lr search.ServiceListResult) (search.ServiceListResult, error) {
				return search.ServiceListResult{}, nil
			},
		),
		nil,
	)

	cl := services.SearchClient{
		Service: m,
	}
	return services.Services{Search: cl}
}

func TestSearchServices(t *testing.T) {
	client.AzureMockTestHelper(t, SearchServices(), buildSearchServices, client.TestOptions{})
}
