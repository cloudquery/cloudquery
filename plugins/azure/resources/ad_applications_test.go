package resources_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildADApplications(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockADApplicationsClient(ctrl)
	var app graphrbac.Application
	faker.SetIgnoreInterface(true)
	defer faker.SetIgnoreInterface(false)
	if err := faker.FakeData(&app); err != nil {
		t.Fatal(err)
	}

	appListPage := graphrbac.NewApplicationListResultPage(
		graphrbac.ApplicationListResult{Value: &[]graphrbac.Application{app}},
		func(ctx context.Context, list graphrbac.ApplicationListResult) (graphrbac.ApplicationListResult, error) {
			return graphrbac.ApplicationListResult{}, nil
		},
	)
	m.EXPECT().List(gomock.Any(), "").Return(appListPage, nil)
	return services.Services{
		AD: services.AD{Applications: m},
	}
}

func TestADApplications(t *testing.T) {
	azureTestHelper(t, resources.AdApplications(), buildADApplications)
}
