package resiliencehub

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildAppVersions(t *testing.T, mock *mocks.MockResiliencehubClient) {
	var l resiliencehub.ListAppVersionsOutput
	if err := faker.FakeObject(&l); err != nil {
		t.Fatal(err)
	}
	l.NextToken = nil
	mock.EXPECT().ListAppVersions(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&l, nil)
	buildAppVersionResources(t, mock)
	buildAppVersionResourceMappings(t, mock)
}
