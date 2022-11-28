package alerts

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/crowdstrike/client"
	"github.com/cloudquery/cloudquery/plugins/source/crowdstrike/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/crowdstrike/gofalcon/falcon/client/alerts"
	"github.com/golang/mock/gomock"
)

func buildQuery(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockAlerts(ctrl)

	var score alerts.GetQueriesAlertsV1OK
	if err := faker.FakeObject(&score); err != nil {
		t.Fatal(err)
	}
	desc := "timestamp.desc"
	mock.EXPECT().GetQueriesAlertsV1(&alerts.GetQueriesAlertsV1Params{
		Context: context.Background(),
		Sort:    &desc,
	}).Return(&score, nil)

	return client.Services{
		Alerts: mock,
	}
}

func TestQuery(t *testing.T) {
	client.MockTestHelper(t, Query(), buildQuery)
}
