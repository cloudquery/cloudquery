package resources_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-11-01-preview/insights"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

// regexMatcher implements gomock.Matcher interface and checks that passed value is a string that matches regular expression in re.
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

func buildActivityLogs(t *testing.T, ctrl *gomock.Controller) services.Services {
	svc := mocks.NewMockActivityLogClient(ctrl)

	ed := insights.EventData{}
	if err := faker.FakeData(&ed); err != nil {
		t.Errorf("failed building mock %s", err)
	}

	filterRe := regexp.MustCompile(`eventTimestamp ge '\d{4}-\d\d-\d\dT\d\d:\d\d:\d\d(\.\d+)Z' and eventTimestamp le '\d{4}-\d\d-\d\dT\d\d:\d\d:\d\d(\.\d+)Z'`)
	svc.EXPECT().List(gomock.Any(), regexMatcher{filterRe}, "").Return(
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
