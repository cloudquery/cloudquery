package resiliencehub

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildSopAlarmRecommendations(t *testing.T, mock *mocks.MockResiliencehubClient) {
	var l resiliencehub.ListSopRecommendationsOutput
	if err := faker.FakeObject(&l); err != nil {
		t.Fatal(err)
	}
	l.NextToken = nil
	mock.EXPECT().ListSopRecommendations(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&l, nil)
}
