package resources_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-11-01-preview/insights"
	resources2 "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildMonitorDiagnosticsSettings(t *testing.T, ctrl *gomock.Controller) services.Services {
	ds := mocks.NewMockDiagnosticSettingsClient(ctrl)
	res := mocks.NewMockResClient(ctrl)

	faker.SetIgnoreInterface(true)

	resource := resources2.GenericResourceExpanded{}
	if err := faker.FakeData(&resource); err != nil {
		t.Fatal(err)
	}
	resourcesPage := resources2.NewListResultPage(
		resources2.ListResult{Value: &[]resources2.GenericResourceExpanded{resource}},
		func(ctx context.Context, result resources2.ListResult) (resources2.ListResult, error) {
			return resources2.ListResult{}, nil
		},
	)
	res.EXPECT().List(gomock.Any(), "", "", nil).Return(resourcesPage, nil)

	d1 := insights.DiagnosticSettingsResource{}
	if err := faker.FakeData(&d1); err != nil {
		t.Fatal(err)
	}
	ds.EXPECT().List(gomock.Any(), "/subscriptions/"+testSubscriptionID).Return(
		insights.DiagnosticSettingsResourceCollection{Value: &[]insights.DiagnosticSettingsResource{d1}}, nil,
	)

	d2 := insights.DiagnosticSettingsResource{}
	if err := faker.FakeData(&d2); err != nil {
		t.Fatal(err)
	}
	ds.EXPECT().List(gomock.Any(), *resource.ID).Return(
		insights.DiagnosticSettingsResourceCollection{Value: &[]insights.DiagnosticSettingsResource{d2}}, nil,
	)

	return services.Services{
		Monitor:   services.MonitorClient{DiagnosticSettings: ds},
		Resources: services.ResourcesClient{Resources: res},
	}
}

func TestMonitorDiagnosticsSettings(t *testing.T) {
	azureTestHelper(t, resources.MonitorDiagnosticSettings(), buildMonitorDiagnosticsSettings)
}
