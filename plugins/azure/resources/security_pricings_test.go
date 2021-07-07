package resources_test

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSecurityPricings(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockSecurityPricingsClient(ctrl)

	pricings := make([]security.Pricing, 0)
	if err := faker.FakeData(&pricings); err != nil {
		t.Fatal(err)
	}

	result := security.PricingList{Value: &pricings}
	m.EXPECT().List(gomock.Any()).Return(result, nil)
	return services.Services{
		Security: services.SecurityClient{Pricings: m},
	}
}

func TestSecurityPricings(t *testing.T) {
	azureTestHelper(t, resources.SecurityPricings(), buildSecurityPricings)
}
