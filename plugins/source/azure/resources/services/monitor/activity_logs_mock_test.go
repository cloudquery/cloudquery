// Auto generated code - DO NOT EDIT.

package monitor

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"regexp"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights"
)

func TestMonitorActivityLogs(t *testing.T) {
	client.MockTestHelper(t, ActivityLogs(), createActivityLogsMock)
}

type regexMatcher struct {
	re *regexp.Regexp
}

func (m regexMatcher) Matches(x interface{}) bool {
	s, ok := x.(string)
	if !ok {
		return false
	}
	return m.re.MatchString(s)
}

func (m regexMatcher) String() string {
	return m.re.String()
}

func createActivityLogsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockMonitorActivityLogsClient(ctrl)
	s := services.Services{
		Monitor: services.MonitorClient{
			ActivityLogs: mockClient,
		},
	}

	data := insights.EventData{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := insights.NewEventDataCollectionPage(insights.EventDataCollection{Value: &[]insights.EventData{data}}, func(ctx context.Context, result insights.EventDataCollection) (insights.EventDataCollection, error) {
		return insights.EventDataCollection{}, nil
	})

	filterRe := regexp.MustCompile(`eventTimestamp ge '\d{4}-\d\d-\d\dT\d\d:\d\d:\d\d(\.\d+)Z' and eventTimestamp le '\d{4}-\d\d-\d\dT\d\d:\d\d:\d\d(\.\d+)Z'`)
	mockClient.EXPECT().List(gomock.Any(), regexMatcher{filterRe}, "").Return(result, nil)
	return s
}
