package resources_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

// nolint:deadcode
func buildADServicePrincipals(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockADServicePrinicpals(ctrl)
	var principal graphrbac.ServicePrincipal
	faker.SetIgnoreInterface(true)
	defer faker.SetIgnoreInterface(false)
	if err := faker.FakeData(&principal); err != nil {
		t.Fatal(err)
	}

	groupListPage := graphrbac.NewServicePrincipalListResultPage(
		graphrbac.ServicePrincipalListResult{Value: &[]graphrbac.ServicePrincipal{principal}},
		func(ctx context.Context, list graphrbac.ServicePrincipalListResult) (graphrbac.ServicePrincipalListResult, error) {
			return graphrbac.ServicePrincipalListResult{}, nil
		},
	)
	m.EXPECT().List(gomock.Any(), "").Return(groupListPage, nil)
	return services.Services{
		AD: services.AD{ServicePrincipals: m},
	}
}

func TestADServicePrincipals(t *testing.T) {
	//azureTestHelper(t, resources.AdServicePrincipals(), buildADServicePrincipals)
}
