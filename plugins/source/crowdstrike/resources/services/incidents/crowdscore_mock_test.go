package incidents

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/crowdstrike/client"
	"github.com/cloudquery/cloudquery/plugins/source/crowdstrike/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/crowdstrike/gofalcon/falcon/client/incidents"
	"github.com/golang/mock/gomock"
)

func buildCrowdScore(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockIncidents(ctrl)

	var score *incidents.CrowdScoreOK
	if err := faker.FakeObject(score); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().CrowdScore(&incidents.CrowdScoreParams{
		// Context: ctx,
		// Sort: &desc,
	}).Return(score, nil)

	return client.Services{
		Incidents: mock,
	}
}

func TestCrowdscore(t *testing.T) {
	client.MockTestHelper(t, Crowdscore(), buildCrowdScore)
}
