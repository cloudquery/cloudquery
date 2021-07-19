package resources_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-11-01-preview/insights"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildActivityLogs(t *testing.T, ctrl *gomock.Controller) services.Services {
	svc := mocks.NewMockActivityLogClient(ctrl)

	ed := insights.EventData{}
	if err := faker.FakeData(&ed); err != nil {
		t.Errorf("failed building mock %s", err)
	}

	svc.EXPECT().List(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		insights.NewEventDataCollectionPage(
			insights.EventDataCollection{Value: &[]insights.EventData{ed}}, func(ctx context.Context, collection insights.EventDataCollection) (insights.EventDataCollection, error) {
				return insights.EventDataCollection{}, nil
			},
		),
		nil,
	)

	s := services.Services{
		Monitor: services.MonitorClient{
			ActivityLogs: svc,
		},
	}
	return s
}

func TestActivityLogs(t *testing.T) {
	azureTestHelper(t, resources.MonitorActivityLogs(), buildActivityLogs)
}
