package resiliencehub

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildAppVersions(t *testing.T, mock *mocks.MockResiliencehubClient) {
	var l resiliencehub.ListAppVersionsOutput
	require.NoError(t, faker.FakeObject(&l))

	l.NextToken = nil
	mock.EXPECT().ListAppVersions(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&l, nil)
	buildAppVersionResources(t, mock)
	buildAppVersionResourceMappings(t, mock)
}
