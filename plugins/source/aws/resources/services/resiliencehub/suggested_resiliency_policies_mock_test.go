package resiliencehub

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildSuggestedResiliencyPolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockResiliencehubClient(ctrl)
	var l resiliencehub.ListSuggestedResiliencyPoliciesOutput
	require.NoError(t, faker.FakeObject(&l))

	l.NextToken = nil
	mock.EXPECT().ListSuggestedResiliencyPolicies(
		gomock.Any(),
		&resiliencehub.ListSuggestedResiliencyPoliciesInput{},
		gomock.Any(),
	).Return(&l, nil)

	return client.Services{Resiliencehub: mock}
}

func TestResiilencehubSuggestedResiliencyPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, SuggestedResiliencyPolicies(), buildSuggestedResiliencyPolicies, client.TestOptions{})
}
